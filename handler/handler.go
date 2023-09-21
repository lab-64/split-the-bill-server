package handler

import (
	"fmt"
	"log"
	"os"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"

	"github.com/google/uuid"

	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/gofiber/fiber/v2"
)

// TODO: Sanitize output / errors
// TODO: Add handler tests

type Handler struct {
	storage storage.AuthenticatedUserStorage
}

func NewHandler(storage storage.AuthenticatedUserStorage) Handler {
	return Handler{storage: storage}
}

// TODO: delete
// tests if we can get the user from the cookie
func (h Handler) TestCookie(c *fiber.Ctx) error {
	cookieID := c.Params("cookie")

	cookieUID, err := uuid.Parse(cookieID)
	// check if given UUID has the right form
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse UUID: %v", err)})
	}
	log.Println(cookieID)
	var user types.User
	user, err = h.storage.GetUserFromAuthCookie(cookieUID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "This cookie does not belong to an user!"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "ok", "user": user})
}

// RegisterUser parses a types.User from the request body, compares and validates both passwords and adds a new user to the storage.
func (h Handler) RegisterUser(c *fiber.Ctx) error {
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}

	// Compare Passwords
	if user.Password != user.ConfirmationPassword {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Bad Password: Password did not match"})
	}
	// Setup Password Validation
	var PasswordValidator = password.NewValidator(false, 8, 64)
	// Load common password list
	var commonPasswords, err = os.Open("common-password-list.txt")
	if err != nil {
		log.Printf("Error while opening file: '%v'", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Server Error: Please try again later"})
	}
	// Validate Passwords
	PasswordValidator.AddCommonPasswords(commonPasswords)
	err = PasswordValidator.ValidatePassword(user.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Bad Password: %v", err)})
	}

	// Add user to storage
	storedUser, err := h.storage.AddUser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create user: %v", err)})
	}

	return c.Status(200).JSON(fiber.Map{"status": "ok", "message": "User successfully created", "user": storedUser.Email})
}

	// TODO: Add cookie to return
	return c.Status(200).JSON(fiber.Map{"status": "ok", "message": "User successfully created", "data": user.Email})
}

// CreateUser parses a types.User from the request body and adds it to the storage.
func (h Handler) CreateUser(c *fiber.Ctx) error {
	log.Printf("CreateUser")
	// Store the body in the user and return error if encountered
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}
	user.ID = uuid.New()
	// Add user to storage.
	_, err := h.storage.AddUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create user: %v", err), "data": err})
	}
	// Return the created user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

func (h Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.storage.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not get users: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": users})
}

// GetUserByID from db
func (h Handler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	user, err := h.storage.GetUserByID(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func (h Handler) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "parameter id is required", "data": nil})
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	err = h.storage.DeleteUser(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to delete user: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func (h Handler) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter username is required", "data": nil})
	}
	user, err := h.storage.GetUserByUsername(username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}
