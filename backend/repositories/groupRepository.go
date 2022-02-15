package repositories

import (
	"go-our-schedule/db"
	"go-our-schedule/dto"
	"go-our-schedule/models"
)

func GetAllMyGroups(id int64) *models.Groups {
	groups := models.Groups{}
	db.DB.Raw("SELECT g.* FROM `go-our-schedule`.groups as g INNER JOIN `go-our-schedule`.members as m ON g.id = m.group_id WHERE user_id = ?", id).
		Scan(&groups)
	return &groups
}

func GetGroup(id int64) (*models.Group, error) {
	group := models.Group{}
	err := db.DB.First(&group, id).Error
	return &group, err
}

func GetGroupByName(name string) (*models.Group, error) {
	group := models.Group{}
	err := db.DB.Where("name = ?", name).First(&group).Error
	return &group, err
}

func CreateGroup(group *models.Group) (*models.Group, error) {
	err := db.DB.Create(group).Error
	return group, err
}

func UpdateGroup(dto *dto.UpdateGroupDto) error {
	return db.DB.Model(models.Group{}).Where("id = ?", dto.Id).
		Updates(map[string]interface{}{"name": dto.Name, "color": dto.Color}).Error
}

func DeleteGroup(id int64) error {
	return db.DB.Delete(&models.Group{}, id).Error
}