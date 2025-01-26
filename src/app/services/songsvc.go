package services

import (
	"EffectiveMobile/src/database/models"
	"EffectiveMobile/src/database/repositories"
)

type SongService struct {
	songRepo  repositories.SongRepository
	groupRepo repositories.GroupRepository
}

func NewSongService(songRepo repositories.SongRepository, groupRepo repositories.GroupRepository) *SongService {
	return &SongService{
		songRepo:  songRepo,
		groupRepo: groupRepo,
	}
}

func (s *SongService) CreateSong(song models.SongRequest) (*models.SongResponse, error) {
	group, err := s.groupRepo.FindGroupByTitle(song.GroupName)
	if err != nil {
		return nil, err
	}

	newSong := models.Song{
		Title:   song.Title,
		GroupID: group.ID,
	}
	if err = s.songRepo.Create(&newSong); err != nil {
		return nil, err
	}

	resSong := s.SongToResponse(&newSong)
	return &resSong, nil
}

func (s *SongService) GetSong(id uint) (*models.SongResponse, error) {
	song, err := s.songRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	resSong := s.SongToResponse(song)
	return &resSong, nil
}

func (s *SongService) GetSongs() ([]models.SongResponse, error) {
	songs, err := s.songRepo.FindAll()
	if err != nil {
		return nil, err
	}

	resSongs := make([]models.SongResponse, len(songs))
	for idx, song := range songs {
		resSongs[idx] = s.SongToResponse(&song)
	}

	return resSongs, err
}

func (s *SongService) SongToResponse(song *models.Song) models.SongResponse {
	group, _ := s.groupRepo.FindByID(song.GroupID)
	return models.SongResponse{
		Id:        song.ID,
		Title:     song.Title,
		GroupName: group.Title,
	}
}
