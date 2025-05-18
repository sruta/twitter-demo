package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"twitter-uala/internal/domain"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/internal/usecase"
	"twitter-uala/pkg"
)

func TestUser_CreateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	inputUser := domain.User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "testpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userMockRepository := repository.NewUserMock()
	userMockRepository.On("SelectByEmail", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	userMockRepository.On("SelectByUsername", mock.Anything).Return(domain.User{}, pkg.NewDBNotFoundError("user", sql.ErrNoRows))
	userMockRepository.On("Insert", mock.Anything).Return(inputUser, nil)
	userMockRepository.On("SelectByID", mock.Anything).Return(inputUser, nil)

	userUsecase := usecase.NewUser(userMockRepository)
	userController := NewUser(userUsecase)

	router := gin.Default()
	router.POST("/user", userController.CreateUser)

	body := map[string]string{
		"username": inputUser.Username,
		"email":    inputUser.Email,
		"password": inputUser.Password,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var created map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &created)
	if err != nil {
		return
	}

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Equal(t, created["id"], float64(inputUser.ID))
	assert.NotNil(t, created["created_at"])
	assert.Nil(t, created["password"])
	userMockRepository.AssertExpectations(t)
}
