package tests

import (
	"EffectiveMobile/src/app/handlers"
	"EffectiveMobile/src/app/services"
	"EffectiveMobile/src/app/validator"
	"EffectiveMobile/tests/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGroup_Success(t *testing.T) {
	mockGroupRepo := new(mocks.MockGroupRepository)
	groupSVC := services.NewGroupService(mockGroupRepo)
	validate := validator.NewValidator()

	mockGroupRepo.On("Create", mock.AnythingOfType("*models.Group")).Return(nil)

	handler := handlers.NewGroupHandler(groupSVC, validate)
	router := gin.Default()
	router.POST("/api/groups", handler.CreateGroup)

	body := bytes.NewBufferString(`{"title": "Test Group"}`)
	req, _ := http.NewRequest("POST", "/api/groups", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"data":{"id":0,"title":"Test Group","songs":[]}`)
	mockGroupRepo.AssertExpectations(t)
}

func TestCreateGroup_Fail(t *testing.T) {
	mockGroupRepo := new(mocks.MockGroupRepository)
	groupSVC := services.NewGroupService(mockGroupRepo)
	validate := validator.NewValidator()

	mockGroupRepo.On("Create", mock.AnythingOfType("*models.Group")).Return(nil)

	handler := handlers.NewGroupHandler(groupSVC, validate)
	router := gin.Default()
	router.POST("/api/groups", handler.CreateGroup)

	body := bytes.NewBufferString(`{"name": "Test Group"}`)
	req, _ := http.NewRequest("POST", "/api/groups", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
