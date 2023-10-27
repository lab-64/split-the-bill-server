package handler

import (
	"fmt"
	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/http"
	"split-the-bill-server/service"
)

type UserHandler struct {
	service.IUserService
	PasswordValidator *password.Validator
}

func NewUserHandler(userService *service.IUserService, v *password.Validator) *UserHandler {
	return &UserHandler{IUserService: *userService, PasswordValidator: v}
}

// GetAll 		func get all users
//
//	@Summary	Get all Users
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.GeneralResponseDTO{data=[]dto.UserOutputDTO}
//	@Router		/api/user [get]
func (h UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.IUserService.GetAll()
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUsersNotFound, err))
	}
	return http.Success(c, fiber.StatusOK, SuccessMsgUsersFound, users)
}

// GetByID 		func get user by id
//
//	@Summary	Get User by ID
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User Id"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.UserOutputDTO}
//	@Router		/api/user/{id} [get]
func (h UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	user, err := h.IUserService.GetByID(uid)
	if err != nil {
		return http.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgUserFound, user)
}

// GetByUsername 	func get user by username
//
//	@Summary	Get User by username
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User Username"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.UserOutputDTO}
//	@Router		/api/user/{username} [get]
func (h UserHandler) GetByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "username"))
	}
	user, err := h.IUserService.GetByUsername(username)
	if err != nil {
		return http.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgUserFound, user)
}

// Create 		func create user
//
//	@Summary	Create User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.UserOutputDTO}
//	@Router		/api/user [post]
func (h UserHandler) Create(c *fiber.Ctx) error {
	// Store the body in the request and return error if encountered
	var request dto.UserInputDTO
	if err := c.BodyParser(&request); err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserParse, err))
	}
	request.ID = uuid.New()
	// Add request to userStorage.
	user, err := h.IUserService.Create(request)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserCreate, err))
	}
	// Return the created request
	return http.Success(c, fiber.StatusOK, SuccessMsgUserCreate, user)
}

// Delete 		func delete user
//
//	@Summary	Delete User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User Username"
//	@Success	200	{object}	dto.GeneralResponseDTO
//	@Router		/api/user/{id} [delete]
func (h UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	err = h.IUserService.Delete(uid)
	if err != nil {
		return http.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserDelete, err))
	}
	return http.Success(c, fiber.StatusOK, SuccessMsgUserDelete, nil)
}

// HandleInvitation 	func handle pending invitation
//
//	@Summary	Handle pending invitation
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.InvitationInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/user/invitations [post]
//
// TODO: Check if id belongs to pending invitation
func (h UserHandler) HandleInvitation(c *fiber.Ctx) error {
	// parse invitation reply
	var request dto.InvitationInputDTO
	if err := c.BodyParser(&request); err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}

	// handle invitation
	err := h.IUserService.HandleInvitation(request)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationHandle, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgInvitationHandle, nil)
}

// Register 	func parses a dto.UserInputDTO from the request body, compares and validates both passwords and adds a new user to the userStorage
//
//	@Summary	Register User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/user/register [post]
func (h UserHandler) Register(c *fiber.Ctx) error {
	var request dto.UserInputDTO
	if err := c.BodyParser(&request); err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserParse, err))
	}

	err := h.PasswordValidator.ValidatePassword(request.Password)
	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBadPassword, err))
	}

	user, err := h.IUserService.Register(request)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserParse, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgUserCreate, user.Username)
}

// Login 		func login user
//
//	@Summary	Login User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CredentialsInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/user/login [post]
//
// Login uses the given login credentials for login and returns an authentication token for the user.
func (h UserHandler) Login(c *fiber.Ctx) error {
	var userCredentials dto.CredentialsInputDTO
	if err := c.BodyParser(&userCredentials); err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserCredentialsParse, err))
	}
	// Checks if all input fields are filled out
	err := userCredentials.ValidateInputs()
	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}

	cookie, err := h.IUserService.Login(userCredentials)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserLogin, err))
	}
	c.Cookie(&cookie)
	return http.Success(c, fiber.StatusOK, SuccessMsgUserLogin, nil)
}
