package impl

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/authentication"
	"split-the-bill-server/common"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
	"split-the-bill-server/storage"
)

type UserService struct {
	storage.IUserStorage
	storage.ICookieStorage
}

func NewUserService(userStorage *storage.IUserStorage, cookieStorage *storage.ICookieStorage) service.IUserService {
	return &UserService{IUserStorage: *userStorage, ICookieStorage: *cookieStorage}
}

func (u *UserService) Create(userDTO dto.UserInputDTO) (dto.UserOutputDTO, error) {
	user := userDTO.ToUser()

	err := u.IUserStorage.Create(user)
	common.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) Delete(id uuid.UUID) error {
	err := u.IUserStorage.Delete(id)
	common.LogError(err)
	return err
}

func (u *UserService) GetAll() ([]dto.UserOutputDTO, error) {
	users, err := u.IUserStorage.GetAll()
	common.LogError(err)

	usersDTO := make([]dto.UserOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = dto.ToUserDTO(&user)
	}

	return usersDTO, err
}

func (u *UserService) GetByID(id uuid.UUID) (dto.UserOutputDTO, error) {
	user, err := u.IUserStorage.GetByID(id)
	common.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) GetByUsername(username string) (dto.UserOutputDTO, error) {
	user, err := u.IUserStorage.GetByUsername(username)
	common.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) Register(userDTO dto.UserInputDTO) (dto.UserOutputDTO, error) {
	user := userDTO.ToUser()
	passwordHash, err := authentication.HashPassword(userDTO.Password)
	common.LogError(err)

	err = u.IUserStorage.Register(user, passwordHash)
	common.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) Login(credentials dto.CredentialsInputDTO) (fiber.Cookie, error) {
	// Log-in user, get authentication cookie
	user, err := u.IUserStorage.GetByUsername(credentials.Username)
	common.LogError(err)

	creds, err := u.IUserStorage.GetCredentials(user.ID)
	common.LogError(err)

	err = authentication.ComparePassword(creds, credentials.Password)
	common.LogError(err)

	sc := authentication.GenerateSessionCookie(user.ID)

	fmt.Printf("%v", user.ID)

	u.ICookieStorage.AddAuthenticationCookie(sc)

	// Create response cookie
	// TODO: add Secure flag after development (cookie will only be sent over HTTPS)
	cookie := fiber.Cookie{
		Name:     authentication.SessionCookieName,
		Value:    sc.Token.String(),
		Expires:  sc.ValidBefore,
		HTTPOnly: true,
		//Secure:   true,
	}

	return cookie, err
}

func (u *UserService) HandleInvitation(invitation dto.InvitationInputDTO) error {
	err := u.IUserStorage.HandleInvitation(invitation.Type, invitation.User, invitation.ID, invitation.Accept)
	common.LogError(err)
	return err
}
