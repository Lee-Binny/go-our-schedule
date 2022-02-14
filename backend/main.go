package main

import (
	"github.com/gin-gonic/gin"
	"go-our-schedule/controllers"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	user := r.Group("/user")
	{
		user.GET("/", controllers.GetAllUsers) // 전체 유저 조회
		user.GET(":id", controllers.GetOneUser) // id로 유저 조회
		user.GET("/search/:name", controllers.GetOneUserByName) // name 유저 조회
		user.POST("/", controllers.SignUp) // 회원가입
		user.POST("/signin", controllers.SignIn) // 로그인
		user.GET("/logout", controllers.Logout) // 로그아웃
	}

	group := r.Group("/group")
	{
		group.GET("/", ) // 전체 그룹 조회
		group.GET("/:id", ) // id로 그룹 조회
		group.POST("/", ) // 그룹 생성
		group.PUT("/", ) // 그룹 정보 수정
		group.DELETE("/") // 그룹 삭제
	}
	
	r.Run()
	return
}
