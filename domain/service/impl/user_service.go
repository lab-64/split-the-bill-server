package impl

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/domain/util"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type UserService struct {
	userStorage   storage.IUserStorage
	cookieStorage storage.ICookieStorage
}

func NewUserService(userStorage *storage.IUserStorage, cookieStorage *storage.ICookieStorage) IUserService {
	return &UserService{userStorage: *userStorage, cookieStorage: *cookieStorage}
}

func (u *UserService) Delete(id uuid.UUID) error {
	err := u.userStorage.Delete(id)
	return err
}

func (u *UserService) GetAll() ([]UserDetailedOutputDTO, error) {
	users, err := u.userStorage.GetAll()
	if err != nil {
		return []UserDetailedOutputDTO{}, err
	}

	usersDTO := make([]UserDetailedOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = ConvertToUserDetailedDTO(&user)
	}

	return usersDTO, err
}

func (u *UserService) GetByID(id uuid.UUID) (UserDetailedOutputDTO, error) {
	user, err := u.userStorage.GetByID(id)
	if err != nil {
		return UserDetailedOutputDTO{}, err
	}

	return ConvertToUserDetailedDTO(&user), err
}

func (u *UserService) Create(userDTO UserInputDTO) (UserCoreOutputDTO, error) {
	user := CreateUserModel(uuid.New(), userDTO)
	passwordHash, err := util.HashPassword(userDTO.Password)
	if err != nil {
		return UserCoreOutputDTO{}, err
	}

	user, err = u.userStorage.Create(user, passwordHash)
	if err != nil {
		return UserCoreOutputDTO{}, err
	}

	return ConvertToUserCoreDTO(&user), err
}

func (u *UserService) Login(credentials CredentialsInputDTO) (UserCoreOutputDTO, AuthCookieModel, error) {
	// Log-in user, get authentication cookie
	user, err := u.userStorage.GetByEmail(credentials.Email)
	if err != nil {
		return UserCoreOutputDTO{}, AuthCookieModel{}, err
	}

	creds, err := u.userStorage.GetCredentials(user.ID)
	if err != nil {
		return UserCoreOutputDTO{}, AuthCookieModel{}, err
	}

	err = util.ComparePassword(creds, credentials.Password)
	if err != nil {
		return UserCoreOutputDTO{}, AuthCookieModel{}, err
	}

	sc := GenerateSessionCookie(user.ID)

	u.cookieStorage.AddAuthenticationCookie(sc)

	return ConvertToUserCoreDTO(&user), sc, err
}
