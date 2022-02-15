package services

import (
	"errors"
	"go-our-schedule/dto"
	"go-our-schedule/models"
	"go-our-schedule/repositories"
)

func GetAllGroups(id int64) *models.Groups {
	return repositories.GetAllMyGroups(id)
}

func GetGroupByName(name string) (*models.Group, error) {
	return repositories.GetGroupByName(name)
}

func CreateGroup(dto *dto.CreateGroupDto) (*models.Group, *models.Member, error) {
	newGroup := models.NewGroup(dto)
	group, err := repositories.CreateGroup(newGroup)
	if err != nil {
		return nil, nil, err
	}

	// TODO 내 id 입력 필요
	newMember := models.CreateMaster(group.Id, 1, dto.UserColor)
	member, err := repositories.CreateMember(newMember)
	if err != nil {
		return nil, nil, err
	}

	return group, member, nil
}

func UpdateGroup(dto *dto.UpdateGroupDto) error {
	_, err := repositories.GetGroup(dto.Id)
	if err != nil {
		return errors.New("not found group")
	}
	// TODO 등급이 Master || Admin 인지 확인
	return repositories.UpdateGroup(dto)
}

func DeleteGroup(id int64) error {
	_, err := repositories.GetGroup(id)
	if err != nil {
		return err
	}
	// TODO 등급이 Master 인지 확인
	err = repositories.DeleteMembers(id)
	if err != nil {
		return err
	}

	return repositories.DeleteGroup(id)
}