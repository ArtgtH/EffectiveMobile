package repositories

import (
	"EffectiveMobile/src/database/models"
	"gorm.io/gorm"
)

type SongRepository interface {
	Create(song *models.Song) error
	FindByID(id uint) (*models.Song, error)
	FindAll() ([]models.Song, error)
	Update(song *models.Song) error
	Delete(id uint) error
}

type SongRepositoryImpl struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepositoryImpl {
	return &SongRepositoryImpl{db: db}
}

func (r *SongRepositoryImpl) Create(song *models.Song) error {
	return r.db.Create(song).Error
}

func (r *SongRepositoryImpl) FindByID(id uint) (*models.Song, error) {
	song := models.Song{}
	if err := r.db.First(&song, id).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *SongRepositoryImpl) FindAll() ([]models.Song, error) {
	var songs []models.Song
	if err := r.db.Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongRepositoryImpl) Update(song *models.Song) error {
	return r.db.Save(song).Error
}

func (r *SongRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Song{}, id).Error
}
