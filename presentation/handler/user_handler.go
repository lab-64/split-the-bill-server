package handler

import (
	"errors"
	"fmt"
	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/domain/util"
	. "split-the-bill-server/presentation"
	"split-the-bill-server/presentation/dto"
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
//	@Success	200	{object}	dto.GeneralResponse{data=[]dto.UserCoreOutput}
//	@Router		/api/user [get]
func (h UserHandler) GetAll(c *fiber.Ctx) error {
	// get all users
	users, err := h.userService.GetAll()
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUsersNotFound, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgUsersFound, users)
}

// GetByID 		func get the user data from a user id
//
//	@Summary	Get User by ID
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User ID"
//	@Success	200	{object}	dto.GeneralResponse{data=dto.UserCoreOutput}
//	@Router		/api/user/{id} [get]
func (h UserHandler) GetByID(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// get user
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
//	@Success	200	{object}	dto.GeneralResponse
//	@Router		/api/user/{id} [delete]
func (h UserHandler) Delete(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// delete user
	err = h.userService.Delete(requesterID, uid)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgUserDelete, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserDelete, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgUserDelete, nil)
}

// Register 	func register user
//
//	@Summary	Register User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserInput	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponse{data=dto.UserCoreOutput}
//	@Router		/api/user [post]
func (h UserHandler) Register(c *fiber.Ctx) error {
	// parse user from request
	var request dto.UserInput
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserParse, err))
	}
	// validate inputs
	err := request.ValidateInputs()
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}
	if err = h.passwordValidator.ValidatePassword(request.Password); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBadPassword, err))
	}
	// create user
	user, err := h.userService.Create(request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgUserCreate, user)
}

// Login 		uses the given login credentials for login and returns an authentication token for the user.
//
//	@Summary	Login User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserInput	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponse{data=dto.UserCoreOutput}
//	@Router		/api/user/login [post]
func (h UserHandler) Login(c *fiber.Ctx) error {
	// parse user from request
	var userCredentials dto.UserInput
	if err := c.BodyParser(&userCredentials); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserCredentialsParse, err))
	}
	// validate inputs
	err := userCredentials.ValidateInputs()
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}
	// login user
	user, sc, err := h.userService.Login(userCredentials)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserLogin, err))
	}
	// create response cookie
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

// Logout 		func logout user
//
//	@Summary	Logout User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.GeneralResponse
//	@Router		/api/user/logout [post]
func (h UserHandler) Logout(c *fiber.Ctx) error {
	// get auth token from request
	token := c.Cookies(middleware.SessionCookieName)
	authToken, err := uuid.Parse(token)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, token, err))
	}
	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// logout user
	err = h.userService.Logout(requesterID, authToken)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserLogout, err))
	}
	// delete cookie
	c.ClearCookie(middleware.SessionCookieName)

	return Success(c, fiber.StatusOK, SuccessMsgUserLogout, nil)
}

// Update 		func update user's username and profile image
//
//	@Summary	Update User
//	@Tags		User
//	@Accept		json
//	@Produce	multipart/form-data
//	@Param		id		path		string			true	"User ID"
//	@Param		request	formData	dto.UserUpdate	true	"Request Body"
//	@Param		image	formData	file			false	"User Image"
//	@Success	200		{object}	dto.GeneralResponse
//	@Router		/api/user/{id} [put]
func (h UserHandler) Update(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// parse user from request
	var user dto.UserUpdate
	if err = c.BodyParser(&user); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserParse, err))
	}
	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// try to parse file
	var content multipart.File
	file, err := c.FormFile("image")
	// TODO: delete
	log.Println("FileName: ", file.Filename)
	// err == nil -> image is included
	if err == nil {
		// read the file content
		content, err = file.Open()
		if err != nil {
			return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserImageUpload, err))
		}
		util.StoreFileInGoogleCloudStorage(content, file.Filename+"BeforeClose")
		defer content.Close()
		// convert file to byte array
		data, fileErr := io.ReadAll(content)
		if fileErr != nil {
			return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserImageUpload, fileErr))
		}
		// check for image type
		contentType := http.DetectContentType(data)
		if err = user.ValidateInputs(contentType); err != nil {
			return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgUserUpdate, err))
		}
		str, err := util.StoreFileInGoogleCloudStorage(content, file.Filename)
		log.Println("-----------: str: ", str, "err ", err)
	}

	// update user
	// TODO: delete
	log.Println("-----------: content: ", content)
	retUser, err := h.userService.Update(requesterID, userID, user, content)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgUserUpdate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserUpdate, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgUserUpdate, retUser)
}
