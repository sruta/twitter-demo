package usecase

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"twitter-uala/internal/domain"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/pkg"
)

func TestUser_Search_Success(t *testing.T) {
	mockRepository := repository.NewUserMock()
	mockRepository.On("Select").Return([]domain.User{}, nil)

	usecase := NewUser(mockRepository)
	users, err := usecase.Search()

	assert.NoError(t, err)
	assert.Empty(t, users)
	mockRepository.AssertExpectations(t)
}

func TestUser_Search_ErrorConn(t *testing.T) {
	mockRepository := repository.NewUserMock()
	mockRepository.On("Select").Return([]domain.User{}, pkg.NewDBFatalError("get users from", sql.ErrConnDone))

	usecase := NewUser(mockRepository)
	users, err := usecase.Search()

	assert.Error(t, err)
	assert.Empty(t, users)
	mockRepository.AssertExpectations(t)
}

func TestUser_SearchByID_Success(t *testing.T) {
	inputID := int64(1)

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByID", mock.Anything).Return(domain.User{ID: inputID}, nil)

	usecase := NewUser(mockRepository)
	user, err := usecase.SearchByID(inputID)

	assert.NoError(t, err)
	assert.Equal(t, inputID, user.ID)
	mockRepository.AssertExpectations(t)
}

func TestUser_SearchByID_ErrorNotFound(t *testing.T) {
	inputID := int64(1)

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByID", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))

	usecase := NewUser(mockRepository)
	user, err := usecase.SearchByID(inputID)

	assert.Error(t, err)
	assert.Equal(t, domain.User{}, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Create_Success(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "testpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByEmail", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("SelectByUsername", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("Insert", mock.Anything).Return(inputUser, nil)
	mockRepository.On("SelectByID", mock.Anything).Return(inputUser, nil)

	usecase := NewUser(mockRepository)
	user, err := usecase.Create(inputUser)

	assert.NoError(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Create_ErrorEmailRepeated(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "testpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByEmail", mock.Anything).Return(inputUser, nil)

	usecase := NewUser(mockRepository)
	user, err := usecase.Create(inputUser)

	assert.Error(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Create_ErrorUsernameRepeated(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "testpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByEmail", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("SelectByUsername", mock.Anything).Return(inputUser, nil)

	usecase := NewUser(mockRepository)
	user, err := usecase.Create(inputUser)

	assert.Error(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Create_ErrorInsertConn(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "testpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByEmail", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("SelectByUsername", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("Insert", mock.Anything).Return(inputUser, pkg.NewDBFatalError("insert user into", sql.ErrConnDone))

	usecase := NewUser(mockRepository)
	user, err := usecase.Create(inputUser)

	assert.Error(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Create_ErrorSelectByIDScan(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "testpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByEmail", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("SelectByUsername", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("Insert", mock.Anything).Return(inputUser, nil)
	mockRepository.On("SelectByID", mock.Anything).Return(inputUser, pkg.NewDBScanFatalError("user", sql.ErrConnDone))

	usecase := NewUser(mockRepository)
	user, err := usecase.Create(inputUser)

	assert.Error(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Update_Success(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser_edited",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dbUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByUsername", inputUser.Username).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	mockRepository.On("SelectByID", inputUser.ID).Return(dbUser, nil).Once()
	dbUser.Username = inputUser.Username
	mockRepository.On("Update", dbUser).Return(inputUser, nil)
	mockRepository.On("SelectByID", inputUser.ID).Return(inputUser, nil).Once()

	usecase := NewUser(mockRepository)
	user, err := usecase.Update(inputUser)

	assert.NoError(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}

func TestUser_Update_ErrorUsernameRepeated(t *testing.T) {
	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dbUser := domain.User{
		ID:        2,
		Email:     "anotheruser@test.com",
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepository := repository.NewUserMock()
	mockRepository.On("SelectByUsername", inputUser.Username).Return(dbUser, nil)

	usecase := NewUser(mockRepository)
	user, err := usecase.Update(inputUser)

	assert.Error(t, err)
	assert.Equal(t, inputUser, user)
	mockRepository.AssertExpectations(t)
}
