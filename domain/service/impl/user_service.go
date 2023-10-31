package impl

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/authentication"
	"split-the-bill-server/core"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service/service_inf"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage/storage_inf"
)

type UserService struct {
	userStorage   storage_inf.IUserStorage
	cookieStorage storage_inf.ICookieStorage
}

func NewUserService(userStorage *storage_inf.IUserStorage, cookieStorage *storage_inf.ICookieStorage) service_inf.IUserService {
	return &UserService{userStorage: *userStorage, cookieStorage: *cookieStorage}
}

func (u *UserService) Create(userDTO dto.UserInputDTO) (dto.UserOutputDTO, error) {
	user := userDTO.ToUser()

	err := u.userStorage.Create(user)
	core.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) Delete(id uuid.UUID) error {
	err := u.userStorage.Delete(id)
	core.LogError(err)
	return err
}

func (u *UserService) GetAll() ([]dto.UserOutputDTO, error) {
	users, err := u.userStorage.GetAll()
	core.LogError(err)

	usersDTO := make([]dto.UserOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = dto.ToUserDTO(&user)
	}

	return usersDTO, err
}

func (u *UserService) GetByID(id uuid.UUID) (dto.UserOutputDTO, error) {
	user, err := u.userStorage.GetByID(id)
	core.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) GetByUsername(username string) (dto.UserOutputDTO, error) {
	user, err := u.userStorage.GetByUsername(username)
	core.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) Register(userDTO dto.UserInputDTO) (dto.UserOutputDTO, error) {
	user := userDTO.ToUser()
	passwordHash, err := authentication.HashPassword(userDTO.Password)
	core.LogError(err)

	err = u.userStorage.Register(user, passwordHash)
	core.LogError(err)

	return dto.ToUserDTO(&user), err
}

func (u *UserService) Login(credentials dto.CredentialsInputDTO) (fiber.Cookie, error) {
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

func (u *UserService) AddGroupInvitation(invitation model.GroupInvitation, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) HandleInvitation(invitation dto.InvitationInputDTO, userID uuid.UUID, invitationID uuid.UUID) error {
	err := u.userStorage.HandleInvitation(invitation.Type, userID, invitationID, invitation.Accept)
	core.LogError(err)
	return err
}

func (u *UserService) GetAuthenticatedUserID(tokenID uuid.UUID) (uuid.UUID, error) {
	// get auth cookie from storage
	cookie, err := u.cookieStorage.GetCookieFromToken(tokenID)
	core.LogError(err)

	// check if cookie is valid
	err = authentication.IsSessionCookieValid(cookie)
	core.LogError(err)

	// get user from cookie
	user, err := u.userStorage.GetByID(cookie.UserID)
	core.LogError(err)

	return user.ID, err
}
