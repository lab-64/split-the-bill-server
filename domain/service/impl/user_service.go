package impl

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	. "github.com/google/uuid"
	"split-the-bill-server/authentication"
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
	return err
}

func (u *UserService) GetAll() ([]UserOutputDTO, error) {
	users, err := u.userStorage.GetAll()
	if err != nil {
		return []UserOutputDTO{}, err
	}

	usersDTO := make([]UserOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = ToUserDTO(&user)
	}

	return usersDTO, err
}

func (u *UserService) GetByID(id UUID) (UserOutputDTO, error) {
	user, err := u.userStorage.GetByID(id)
	if err != nil {
		return UserOutputDTO{}, err
	}

	return ToUserDTO(&user), err
}

func (u *UserService) Create(userDTO UserInputDTO) (UserOutputDTO, error) {
	user := ToUserModel(userDTO)
	passwordHash, err := authentication.HashPassword(userDTO.Password)
	if err != nil {
		return UserOutputDTO{}, err
	}

	user, err = u.userStorage.Create(user, passwordHash)
	if err != nil {
		return UserOutputDTO{}, err
	}

	return ToUserDTO(&user), err
}

func (u *UserService) Login(credentials CredentialsInputDTO) (UserOutputDTO, fiber.Cookie, error) {
	// Log-in user, get authentication cookie
	user, err := u.userStorage.GetByEmail(credentials.Email)
	if err != nil {
		return UserOutputDTO{}, fiber.Cookie{}, err
	}

	creds, err := u.userStorage.GetCredentials(user.ID)
	if err != nil {
		return UserOutputDTO{}, fiber.Cookie{}, err
	}

	err = authentication.ComparePassword(creds, credentials.Password)
	if err != nil {
		return UserOutputDTO{}, fiber.Cookie{}, err
	}

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

	return ToUserDTO(&user), cookie, err
}
