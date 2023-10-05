package handler

import (
	"fmt"
	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/authentication"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"split-the-bill-server/wire"

	"github.com/gofiber/fiber/v2"
)

// TODO: Sanitize output / errors
// TODO: Add handler tests

type Handler struct {
	userStorage       storage.UserStorage
	cookieStorage     storage.CookieStorage
	PasswordValidator *password.Validator
}

func NewHandler(userStorage storage.UserStorage, cookieStorage storage.CookieStorage, v *password.Validator) Handler {
	return Handler{userStorage: userStorage, cookieStorage: cookieStorage, PasswordValidator: v}
}

// RegisterUser parses a types.User from the request body, compares and validates both passwords and adds a new user to the userStorage.
func (h Handler) RegisterUser(c *fiber.Ctx) error {
	var rUser wire.RegisterUser
	if err := c.BodyParser(&rUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse rUser: %v", err), "data": err})
	}

	err := h.PasswordValidator.ValidatePassword(rUser.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Bad Password: %v", err)})
	}

	user := rUser.ToUser()
	pHash, err := authentication.HashPassword(rUser.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create user")})

	}

	err = h.userStorage.RegisterUser(user, pHash)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create rUser: %v", err)})
	}

	return c.Status(200).JSON(fiber.Map{"status": "ok", "message": "User successfully created", "rUser": user.Email})
}

// Login uses the given login credentials for login and returns an authentication token for the user.
func (h Handler) Login(c *fiber.Ctx) error {
	var userCredentials wire.Credentials
	if err := c.BodyParser(&userCredentials); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}
	// Checks if all input fields are filled out
	err := userCredentials.ValidateInputs()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Inputs invalid: %v", err)})
	}
	// Log-in user, get authentication cookie
	user, err := h.userStorage.GetUserByUsername(userCredentials.Username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}
	creds, err := h.userStorage.GetCredentials(user.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}
	err = authentication.ComparePassword(creds, userCredentials.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}

	sc := authentication.GenerateSessionCookie(user.ID)

	h.cookieStorage.AddAuthenticationCookie(sc)

	return c.Status(200).JSON(fiber.Map{"status": "ok", "cookieAuth": sc})
}

// CreateUser parses a types.User from the request body and adds it to the userStorage.
func (h Handler) CreateUser(c *fiber.Ctx) error {
	log.Println("CreateUser")
	// Store the body in the user and return error if encountered
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}
	user.ID = uuid.New()
	// Add user to userStorage.
	err := h.userStorage.AddUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create user: %v", err), "data": err})
	}
	// Return the created user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

func (h Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userStorage.GetAllUsers()
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
	user, err := h.userStorage.GetUserByID(uid)
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
	err = h.userStorage.DeleteUser(uid)
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
	user, err := h.userStorage.GetUserByUsername(username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}
