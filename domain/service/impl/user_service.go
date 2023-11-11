package impl

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	. "github.com/google/uuid"
	"split-the-bill-server/authentication"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
	. "split-the-bill-server/storage/storage_inf"
)

type UserService struct {
	userStorage   IUserStorage
	cookieStorage ICookieStorage
}

func NewUserService(userStorage *IUserStorage, cookieStorage *ICookieStorage) IUserService {
	return &UserService{userStorage: *userStorage, cookieStorage: *cookieStorage}
}

func (u *UserService) Delete(id UUID) error {
	err := u.userStorage.Delete(id)
	core.LogError(err)
	return err
}

func (u *UserService) GetAll() ([]UserOutputDTO, error) {
	users, err := u.userStorage.GetAll()
	core.LogError(err)

	usersDTO := make([]UserOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = ToUserDTO(&user)
	}

	return usersDTO, err
}

func (u *UserService) GetByID(id UUID) (UserOutputDTO, error) {
	user, err := u.userStorage.GetByID(id)
	core.LogError(err)

	return ToUserDTO(&user), err
}

func (u *UserService) GetByUsername(username string) (UserOutputDTO, error) {
	user, err := u.userStorage.GetByUsername(username)
	core.LogError(err)

	return ToUserDTO(&user), err
}

func (u *UserService) Register(userDTO UserInputDTO) (UserOutputDTO, error) {
	user := ToUserModel(userDTO)
	passwordHash, err := authentication.HashPassword(userDTO.Password)
	core.LogError(err)

	err = u.userStorage.Create(user, passwordHash)
	core.LogError(err)

	return ToUserDTO(&user), err
}

func (u *UserService) Login(credentials CredentialsInputDTO) (fiber.Cookie, error) {
	// Log-in user, get authentication cookie
	user, err := u.userStorage.GetByUsername(credentials.Username)
	core.LogError(err)

	creds, err := u.userStorage.GetCredentials(user.ID)
	core.LogError(err)

	err = authentication.ComparePassword(creds, credentials.Password)
	core.LogError(err)

	sc := authentication.GenerateSessionCookie(user.ID)

	fmt.Printf("%v", user.ID)

	u.cookieStorage.AddAuthenticationCookie(sc)

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

func (u *UserService) HandleInvitation(invitation InvitationInputDTO) error {
	err := u.userStorage.HandleInvitation(invitation.Type, invitation.User, invitation.ID, invitation.Accept)
	core.LogError(err)
	return err
}
