package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/converter"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/domain/util"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
	"strings"
)

type UserService struct {
	userStorage   storage.IUserStorage
	cookieStorage storage.ICookieStorage
}

func NewUserService(userStorage *storage.IUserStorage, cookieStorage *storage.ICookieStorage) IUserService {
	return &UserService{userStorage: *userStorage, cookieStorage: *cookieStorage}
}

func (u *UserService) Delete(requesterID uuid.UUID, id uuid.UUID) error {
	// Authorization
	if requesterID != id {
		return domain.ErrNotAuthorized
	}

	err := u.userStorage.Delete(id)
	return err
}

func (u *UserService) GetAll() ([]dto.UserCoreOutput, error) {
	users, err := u.userStorage.GetAll()
	if err != nil {
		return []dto.UserCoreOutput{}, err
	}

	usersDTO := make([]dto.UserCoreOutput, len(users))

	for i, user := range users {
		usersDTO[i] = converter.ToUserCoreDTO(&user)
	}

	return usersDTO, err
}

func (u *UserService) GetByID(id uuid.UUID) (dto.UserCoreOutput, error) {
	user, err := u.userStorage.GetByID(id)
	if err != nil {
		return dto.UserCoreOutput{}, err
	}

	return converter.ToUserCoreDTO(&user), err
}

func (u *UserService) Create(userDTO dto.UserInput) (dto.UserCoreOutput, error) {
	// extract username from email
	username := strings.Split(userDTO.Email, "@")[0]
	user := model.CreateUser(uuid.New(), userDTO.Email, username, "")
	passwordHash, err := util.HashPassword(userDTO.Password)
	if err != nil {
		return dto.UserCoreOutput{}, err
	}

	user, err = u.userStorage.Create(user, passwordHash)
	if err != nil {
		return dto.UserCoreOutput{}, err
	}

	return converter.ToUserCoreDTO(&user), err
}

func (u *UserService) Login(userInput dto.UserInput) (dto.UserCoreOutput, model.AuthCookie, error) {
	// Log-in user, get authentication cookie
	user, err := u.userStorage.GetByEmail(userInput.Email)
	if err != nil {
		return dto.UserCoreOutput{}, model.AuthCookie{}, err
	}

	credentials, err := u.userStorage.GetCredentials(user.ID)
	if err != nil {
		return dto.UserCoreOutput{}, model.AuthCookie{}, err
	}

	err = util.ComparePassword(credentials, userInput.Password)
	if err != nil {
		return dto.UserCoreOutput{}, model.AuthCookie{}, err
	}

	sc := model.GenerateSessionCookie(user.ID)

	u.cookieStorage.AddAuthenticationCookie(sc)

	return converter.ToUserCoreDTO(&user), sc, err
}

func (u *UserService) Update(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate, file []byte) (dto.UserCoreOutput, error) {
	// Authorization
	if requesterID != id {
		return dto.UserCoreOutput{}, domain.ErrNotAuthorized
	}
	// Get user
	userModel, err := u.userStorage.GetByID(id)
	if err != nil {
		return dto.UserCoreOutput{}, err
	}
	// store profile image if file is included
	filePath := ""
	if file != nil {
		filePath, err = util.StoreFile(file, id)
		if err != nil {
			return dto.UserCoreOutput{}, err
		}
	} else { // if no file is included, use the old path
		filePath = userModel.ProfileImgPath
	}
	// update user's username and profile image
	userModel.Username = user.Username
	userModel.ProfileImgPath = filePath
	updatedUser, err := u.userStorage.Update(userModel)
	if err != nil {
		return dto.UserCoreOutput{}, err
	}

	return converter.ToUserCoreDTO(&updatedUser), err
}
