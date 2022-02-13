package repositories

import (
	"go-our-schedule/db"
	"go-our-schedule/models"
)

func GetAllUsers() *models.Users {
	users := models.Users{}
	db.DB.Find(&users)
	return &users
}

func GetOneUser(id int64) (*models.User, error) {
	user := models.User{}
	err := db.DB.First(&user, id).Error
	return &user, err
}

func GetOneUserByUid(uid string) (*models.User, error) {
	user := models.User{}
	err := db.DB.Where("uid = ?", uid).First(&user).Error
	return &user, err
}

func GetOneUserByName(name string) (*models.Users, error) {
	users := models.Users{}
	err := db.DB.Where("name = ?", name).Find(&users).Error
	return &users, err
}

func GetUsersName(id []int64) (*models.Users, error) {
	users := models.Users{}
	err := db.DB.Select("name").Find(&users, id).Error
	return &users, err
}

func CreateUser(user *models.User) (*models.User, error) {
	err := db.DB.Create(user).Error
	return user, err
}