package services

import (
	"EffectiveMobile/src/database/models"
	"EffectiveMobile/src/database/repositories"
)

type GroupService struct {
	repo repositories.GroupRepository
}

func NewGroupService(repo repositories.GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(group models.GroupRequest) (*models.GroupResponse, error) {
	newGroup := models.Group{
		Title: group.Title,
		Songs: []models.Song{},
	}
	if err := s.repo.Create(&newGroup); err != nil {
		return nil, err
	}

	resGroup := GroupToResponse(&newGroup)
	return &resGroup, nil
}

func (s *GroupService) GetGroup(id uint) (*models.GroupResponse, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	resGroup := GroupToResponse(group)
	return &resGroup, nil
}

func (s *GroupService) GetGroups() ([]models.GroupResponse, error) {
	groups, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	resGroups := make([]models.GroupResponse, len(groups))
	for idx, group := range groups {
		resGroups[idx] = GroupToResponse(&group)
	}

	return resGroups, err
}

func GroupToResponse(group *models.Group) models.GroupResponse {
	return models.GroupResponse{
		Id:    group.ID,
		Title: group.Title,
		Songs: group.Songs,
	}
}
