package controllers

import (
	"github.com/gin-gonic/gin"
	"go-our-schedule/dto"
	"go-our-schedule/services"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": services.GetAllUsers(),
	})
}

func GetOneUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	user, err := services.GetOneUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetOneUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := services.GetOneUserByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func SignUp(c *gin.Context) {
	req := &dto.SignUpUserDto{}
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "invalid input data",
		})
		return
	}

	user, err := services.SignUp(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func SignIn(c *gin.Context) {
	req := &dto.SignInUserDto{}
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "invalid input data",
		})
		return
	}

	// TODO 토큰이나 세션 추가
	user, err := services.SignIn(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Logout(c *gin.Context) {
	// TODO 토큰이나 세션 삭제
}