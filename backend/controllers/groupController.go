package controllers

import (
	"github.com/gin-gonic/gin"
	"go-our-schedule/dto"
	"go-our-schedule/services"
	"net/http"
	"strconv"
)

func GetAllMyGroups(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"groups": services.GetAllGroups(1),
	})
}

func GetGroupByName(c *gin.Context) {
	name := c.Param("name")
	group, err := services.GetGroupByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"group": group,
	})
}

func CreateGroup(c *gin.Context) {
	req := &dto.CreateGroupDto{}
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"message": "invalid input data",
		})
		return
	}

	group, member, err := services.CreateGroup(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"group": group,
		"member": member,
	})
}

func UpdateGroup(c *gin.Context) {
	req := &dto.UpdateGroupDto{}
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": false,
			"message": "invalid input data",
		})
		return
	}

	err = services.UpdateGroup(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}

func DeleteGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"message": "invalid id",
		})
		return
	}

	if err = services.DeleteGroup(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}