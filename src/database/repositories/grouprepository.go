package repositories

import (
	"EffectiveMobile/src/database/models"
	"gorm.io/gorm"
)

type GroupRepository interface {
	Create(group *models.Group) error
	FindByID(id uint) (*models.Group, error)
	FindAll() ([]models.Group, error)
	Update(group *models.Group) error
	Delete(id uint) error
	FindGroupByTitle(title string) (*models.Group, error)
}

type GroupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &GroupRepositoryImpl{db: db}
}

func (r *GroupRepositoryImpl) Create(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *GroupRepositoryImpl) FindByID(id uint) (*models.Group, error) {
	group := models.Group{}
	if err := r.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepositoryImpl) FindAll() ([]models.Group, error) {
	var groups []models.Group
	if err := r.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *GroupRepositoryImpl) Update(group *models.Group) error {
	return r.db.Save(group).Error
}

func (r *GroupRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Group{}, id).Error
}

func (r *GroupRepositoryImpl) FindGroupByTitle(title string) (*models.Group, error) {
	group := models.Group{}
	if err := r.db.First(&group, "title = ?", title).Error; err != nil {
		return nil, err
	}
	return &group, nil
}
