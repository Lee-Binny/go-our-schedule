package repositories

import (
	"go-our-schedule/db"
	"go-our-schedule/models"
)

func CreateMember(member *models.Member) (*models.Member, error) {
	err := db.DB.Create(member).Error
	return member, err
}

func UpdateMemberGrade(groupId, userId int64, grade int) error {
	return db.DB.Model(models.Member{}).Where("group_id = ? And user_id", groupId, userId).
		Updates(map[string]interface{}{"grade": grade}).Error
}

func DeleteMembers(groupId int64) error {
	return db.DB.Delete(&models.Group{}).Where("group_id = ?", groupId).Error
}