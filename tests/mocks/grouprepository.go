package mocks

import (
	"EffectiveMobile/src/database/models"
	"github.com/stretchr/testify/mock"
)

type MockGroupRepository struct {
	mock.Mock
}

func (m *MockGroupRepository) Create(group *models.Group) error {
	args := m.Called(group)
	return args.Error(0)
}

func (m *MockGroupRepository) FindByID(id uint) (*models.Group, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Group), args.Error(1)
}

func (m *MockGroupRepository) FindAll() ([]models.Group, error) {
	args := m.Called()
	return args.Get(0).([]models.Group), args.Error(1)
}

func (m *MockGroupRepository) Update(group *models.Group) error {
	args := m.Called(group)
	return args.Error(0)
}

func (m *MockGroupRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockGroupRepository) FindGroupByTitle(title string) (*models.Group, error) {
	args := m.Called(title)
	return args.Get(0).(*models.Group), args.Error(1)
}
