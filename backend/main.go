package main

import (
	"github.com/gin-gonic/gin"
	"go-our-schedule/db"
	"log"
)

func main()  {
	r := gin.Default()

	err := db.ConnectDB()
	if err != nil {
		log.Fatal("database connection error", err)
		return
	}

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "hello world",
	//	})
	//})
	r.Run()
	return
}
