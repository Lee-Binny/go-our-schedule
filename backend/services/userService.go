package services

import (
	"github.com/gin-gonic/gin"
	"go-our-schedule/models"
	"go-our-schedule/repositories"
)

func GetAllUsers(c *gin.Context) *models.Users{
	return repositories.GetAllUsers()
}

func GetOneUser(c *gin.Context) {

}

func GetOneUserByName(c *gin.Context) {

}

func SignUp(c *gin.Context) {

}

func SignIn(c *gin.Context) {

}

func Logout(c *gin.Context) {

}