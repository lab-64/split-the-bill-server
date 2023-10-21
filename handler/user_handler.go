package handler

import (
	"errors"
	"fmt"
	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
)

type UserHandler struct {
	service.IUserService
	service.ICookieService
	PasswordValidator *password.Validator
}

func NewUserHandler(userService *service.IUserService, cookieService *service.ICookieService, v *password.Validator) *UserHandler {
	return &UserHandler{IUserService: *userService, ICookieService: *cookieService, PasswordValidator: v}
}

func (h UserHandler) Route(api fiber.Router) {
	user := api.Group("/user")

	// Get all Users
	user.Get("/", h.GetAll)

	// Get User by ID
	user.Get("/:id", h.GetByID)

	// Get User by Username
	user.Get("/:username", h.GetByUsername)

	// Create User
	user.Post("/", h.Create)

	// Register User
	user.Post("/register", h.Register)

	// Login User
	user.Post("/login", h.Login)

	// Update User
	//user.Put("/:id", h.UpdateUser)

	// Delete User by ID
	user.Delete("/:id", h.Delete)

	// Handle invitation //TODO: this endpoint might be too complex and not intuitive, maybe split it up
	user.Post("/invitations", h.HandleInvitation)
}

func (h UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.IUserService.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not get users: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": users})
}

func (h UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	user, err := h.IUserService.GetByID(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func (h UserHandler) GetByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter username is required", "data": nil})
	}
	user, err := h.IUserService.GetByUsername(username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// Create parses a types.User from the request body and adds it to the userStorage.
func (h UserHandler) Create(c *fiber.Ctx) error {
	// Store the body in the request and return error if encountered
	var request dto.UserCreateDTO
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse request: %v", err), "data": err})
	}
	request.ID = uuid.New()
	// Add request to userStorage.
	user, err := h.IUserService.Create(request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create request: %v", err), "data": err})
	}
	// Return the created request
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

func (h UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "parameter id is required", "data": nil})
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	err = h.IUserService.Delete(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to delete user: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

// TODO: Check if id belongs to pending invitation
func (h UserHandler) HandleInvitation(c *fiber.Ctx) error {
	// get authenticated user
	userID, err := h.getAuthenticatedUserFromHeader(c.GetReqHeaders())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Authentication declined: %v", err)})
	}
	// parse invitation reply
	var request dto.InvitationReplyDTO
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse invitation: %v", err), "data": err})
	}

	// handle invitation
	err = h.IUserService.HandleInvitation(request, userID, request.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not handle invitation: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Invitation handled"})
}

// Register parses a types.User from the request body, compares and validates both passwords and adds a new user to the userStorage.
func (h UserHandler) Register(c *fiber.Ctx) error {
	var request dto.UserCreateDTO
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse request: %v", err), "data": err})
	}

	err := h.PasswordValidator.ValidatePassword(request.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Bad Password: %v", err)})
	}

	user, err := h.IUserService.Register(request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create User: %v", err)})
	}

	return c.Status(200).JSON(fiber.Map{"status": "ok", "message": "User successfully created", "User": user.Username})
}

// Login uses the given login credentials for login and returns an authentication token for the user.
func (h UserHandler) Login(c *fiber.Ctx) error {
	var userCredentials dto.CredentialsDTO
	if err := c.BodyParser(&userCredentials); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}
	// Checks if all input fields are filled out
	err := userCredentials.ValidateInputs()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Inputs invalid: %v", err)})
	}

	err = h.IUserService.Login(c, userCredentials)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}

	return c.Status(200).JSON(fiber.Map{"status": "ok"})
}

// GetAuthenticatedUserFromHeader tries to return the user id associated with the given authentication token in the request header.
// If the token is invalid, an error will be returned.
// TODO: Generalize error messages
func (h UserHandler) getAuthenticatedUserFromHeader(reqHeader map[string]string) (uuid.UUID, error) {
	// get authentication cookie from header
	token := reqHeader["Cookie"]
	// check if cookie is present
	if token == "" {
		return uuid.Nil, errors.New("authentication cookie is missing")
	}
	// try to parse token
	tokenUUID, err := uuid.Parse(token)
	if err != nil {
		return uuid.Nil, errors.New("authentication cookie is invalid")
	}

	userID, err := h.IUserService.GetAuthenticatedUserID(tokenUUID)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}
