package handler

import (
	"errors"
	"fmt"
	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path/filepath"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/middleware"
)

type UserHandler struct {
	userService       service.IUserService
	passwordValidator *password.Validator
}

func NewUserHandler(userService *service.IUserService, v *password.Validator) *UserHandler {
	return &UserHandler{userService: *userService, passwordValidator: v}
}

// GetAll 		func get all users
//
//	@Summary	Get all Users
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.GeneralResponseDTO{data=[]dto.UserDetailedOutputDTO}
//	@Router		/api/user [get]
func (h UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUsersNotFound, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgUsersFound, users)
}

// GetByID 		func get the detailed user data from a user id
//
//	@Summary	Get detailed User data by ID
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.UserDetailedOutputDTO}
//	@Router		/api/user/{id} [get]
func (h UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	user, err := h.userService.GetByID(uid)
	if err != nil {
		return Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgUserFound, user)
}

// Delete 		func delete user
//
//	@Summary	Delete User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User ID"
//	@Success	200	{object}	dto.GeneralResponseDTO
//	@Router		/api/user/{id} [delete]
func (h UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	err = h.userService.Delete(uid)
	if err != nil {
		return Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserDelete, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgUserDelete, nil)
}

// Register 	parses a dto.UserInputDTO from the request body, compares and validates both passwords and creates a new user.
//
//	@Summary	Register User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.UserCoreOutputDTO}
//	@Router		/api/user [post]
func (h UserHandler) Register(c *fiber.Ctx) error {
	var request UserInputDTO
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserParse, err))
	}

	if err := h.passwordValidator.ValidatePassword(request.Password); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBadPassword, err))
	}

	user, err := h.userService.Create(request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgUserCreate, user)
}

// Login 		func login user
//
//	@Summary	Login User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CredentialsInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.UserCoreOutputDTO}
//	@Router		/api/user/login [post]
//
// Login uses the given login credentials for login and returns an authentication token for the user.
func (h UserHandler) Login(c *fiber.Ctx) error {
	var userCredentials CredentialsInputDTO
	if err := c.BodyParser(&userCredentials); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserCredentialsParse, err))
	}
	// Checks if all input fields are filled out
	err := userCredentials.ValidateInputs()
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}

	user, sc, err := h.userService.Login(userCredentials)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserLogin, err))
	}

	// Create response cookie
	// TODO: add Secure flag after development (cookie will only be sent over HTTPS)
	cookie := fiber.Cookie{
		Name:     middleware.SessionCookieName,
		Value:    sc.Token.String(),
		Expires:  sc.ValidBefore,
		HTTPOnly: true,
		//Secure:   true,
	}

	c.Cookie(&cookie)
	return Success(c, fiber.StatusOK, SuccessMsgUserLogin, user)
}

// Update 		func update user
//
//	@Summary	Update User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"User ID"
//	@Param		request	body		dto.UserUpdateDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/user/{id} [put]
func (h UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	var user UserUpdateDTO
	if err := c.BodyParser(&user); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserParse, err))
	}

	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)

	retUser, err := h.userService.Update(requesterID, userID, user)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgUserUpdate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserUpdate, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgUserUpdate, retUser)
}

// UploadImage 		func upload user image
//
//	@Summary	Upload User Image
//	@Tags		User
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		id		path		string	true	"User ID"
//	@Param		image	formData	file	true	"User Image"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/user/upload/{id} [post]
func (h UserHandler) UploadImage(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	_, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// get file
	file, err := c.FormFile("image")
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserImageUpload, err))
	}
	// convert file to byte array
	// Read the file content
	content, err := file.Open()
	if err != nil {
		return err
	}
	defer content.Close()
	data, err := io.ReadAll(content)
	if err != nil {
		return err
	}

	storagePath := "./uploads/" + id

	// create storage directory for id
	if err = os.MkdirAll(storagePath, os.ModePerm); err != nil {
		return err
	}

	// save file
	filePath := filepath.Join(storagePath, file.Filename)
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	log.Println("File saved to: " + filePath)

	// Read the image file
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Set the appropriate content type for the image
	c.Set("Content-Type", "image/jpeg")

	// Return the image data in the response body
	return c.Send(imageData)
}
