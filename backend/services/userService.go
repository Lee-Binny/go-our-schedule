package services

import (
	"errors"
	"go-our-schedule/dto"
	"go-our-schedule/models"
	"go-our-schedule/repositories"
)

func GetAllUsers() *models.Users{
	return repositories.GetAllUsers()
}

func GetOneUser(id int64) (*models.User, error) {
	return repositories.GetOneUser(id)
}

func GetOneUserByName(name string) (*models.User, error) {
	return repositories.GetOneUserByName(name)
}

func SignUp(dto *dto.SignUpUserDto) (*models.User, error) {
	newUser := models.NewUser(dto)
	err := newUser.HashPassword()
	if err != nil {
		return nil, errors.New("not hashed password")
	}

	return repositories.CreateUser(newUser)
}

func SignIn(dto *dto.SignInUserDto) (*models.User, error) {
	user, err := repositories.GetOneUserByUid(dto.UserId)
	if err != nil {
		return nil, err
	}

	err = user.IsMatchedPassword(dto.Password)
	if err != nil {
		return nil, errors.New("not matched password")
	}

	return user, nil
}

func Logout() {

}