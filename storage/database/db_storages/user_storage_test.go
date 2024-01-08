package db_storages

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"split-the-bill-server/authentication"
	. "split-the-bill-server/domain/service/impl"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/handler"
	. "split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/entity"
	. "split-the-bill-server/storage/database/test_util"
	"testing"
)

func TestCase(t *testing.T) {
	db, err := database.NewDatabase()
	if err != nil {
		t.Fatal(err)
	}

	passwordValidator, err := authentication.NewPasswordValidator()
	if err != nil {
		t.Fatal(err)
	}

	userStorage := NewUserStorage(db)
	cookieStorage := NewCookieStorage(db)
	userService := NewUserService(&userStorage, &cookieStorage)
	userHandler := handler.NewUserHandler(&userService, passwordValidator)
	log.Println(userHandler)

	// Test case: Successful creation
	user := dto.UserInputDTO{
		Email:    "felix@mail.com",
		Password: "alek1337",
	}
	log.Println(user)
}

func TestUserStorage_Create_Success(t *testing.T) {
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()

	// Create an instance of UserStorage with the mocked DB
	userStorage := UserStorage{DB: gormDB}

	currentUser := User

	// Expectations for the transaction with mocked behavior
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(currentUser.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), currentUser.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO "credentials"`).
		WithArgs(currentUser.ID, PasswordHash).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userStorage.Create(currentUser, PasswordHash)
	assert.NoError(t, err) // Check if the Create method returns no error

	// Ensure all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUserStorage_Create_Bad_Inputs(t *testing.T) {
	// Initialize mock DB and UserStorage
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()
	userStorage := UserStorage{DB: gormDB}

	// Test case: Create user with empty email
	userWithEmptyEmail := UserWithEmptyEmail
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userWithEmptyEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), userWithEmptyEmail.Email).
		WillReturnError(InvalidUserInputError)
	mock.ExpectRollback()

	err := userStorage.Create(userWithEmptyEmail, PasswordHash)
	assert.Error(t, err)
	assert.EqualError(t, err, InvalidUserInputError.Error())

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUserStorage_Create_Already_Exist(t *testing.T) {
	// Initialize mock DB and UserStorage
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()
	userStorage := UserStorage{DB: gormDB}

	// Test case: Successful creation
	user := User
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(user.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO "credentials"`).
		WithArgs(user.ID, PasswordHash).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userStorage.Create(user, PasswordHash)
	assert.NoError(t, err)

	// Test case: Create user with same ID
	userWithSameID := UserWithSameID
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userWithSameID.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), userWithSameID.Email).
		WillReturnError(errors.New("duplicate key value violates unique constraint \"users_pkey\""))
	mock.ExpectRollback()

	err = userStorage.Create(userWithSameID, PasswordHash)
	assert.Error(t, err)
	assert.EqualError(t, err, UserAlreadyExistsError.Error())

	// Test case: Create user with same email
	userWithSameEmail := UserWithSameEmail
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userWithSameEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), userWithSameEmail.Email).
		WillReturnError(errors.New("duplicate key value violates unique constraint \"users_email_key\""))
	mock.ExpectRollback()

	err = userStorage.Create(userWithSameEmail, PasswordHash)
	assert.Error(t, err)
	assert.EqualError(t, err, UserAlreadyExistsError.Error())

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetByID(t *testing.T) {
	// Initialize mock DB and UserStorage
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()
	userStorage := UserStorage{DB: gormDB}

	// test data
	userIDSuccess := uuid.New()

	tests := []struct {
		name    string
		storage UserStorage
		userUID uuid.UUID
		mock    func()
		wantErr bool
		want    entity.User
	}{
		{
			// When everything works as expected
			name:    "Get User by ID Success",
			storage: userStorage,
			userUID: userIDSuccess,
			mock: func() {
				// ExpectQuery for SELECT Query
				// ExpectExec for INSERT, UPDATE, DELETE, ...
				// We added one row
				userRows := sqlmock.NewRows([]string{"ID", "Email"}).AddRow(userIDSuccess, "mail@mail.com")
				mock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(userIDSuccess).WillReturnRows(userRows)
				groupInvitationRows := sqlmock.NewRows([]string{"ID", "InviteeID"}).AddRow(uuid.New(), userIDSuccess) // Include field where user is stored
				mock.ExpectQuery(`SELECT (.+) FROM "group_invitations"`).WithArgs(userIDSuccess).WillReturnRows(groupInvitationRows)
				groupMemberRows := sqlmock.NewRows([]string{"ID", "OwnerUID"}).AddRow(uuid.New(), userIDSuccess)
				mock.ExpectQuery(`SELECT (.+) FROM "group_members"`).WithArgs(userIDSuccess).WillReturnRows(groupMemberRows)
			},
			want: entity.User{Base: entity.Base{ID: userIDSuccess}},
		},
		/*		{
					//When the role tried to access is not found
					name:    "Not Found",
					storage:       userStorage,
					userUID: uuid.New(),
					mock: func() {
						rows := sqlmock.NewRows([]string{"Id", "Email", "CreatedAt"}) //observe that we didnt add any role here
						mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(1).WillReturnRows(rows)
					},
					wantErr: true,
					want:    entity.User{},
				},
				{
					//When invalid statement is provided, ie the SQL syntax is wrong(in this case, we provided a wrong database)
					name:    "Invalid Prepare",
					storage:       userStorage,
					userUID: uuid.New(),
					mock: func() {
						rows := sqlmock.NewRows([]string{"Id", "Title", "Body", "CreatedAt"}).AddRow(1, "title", "body", created_at)
						mock.ExpectPrepare("SELECT (.+) FROM wrong_table").ExpectQuery().WithArgs(1).WillReturnRows(rows)
					},
					wantErr: true,
					want:    entity.User{},
				},*/
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			got, err := testcase.storage.GetByID(testcase.userUID)
			if (err != nil) != testcase.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, testcase.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got.ID, testcase.want.ID) {
				t.Errorf("Get() = %v, want %v", got, testcase.want)
			}
		})
	}
}
