package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hikobend/spl/controller"
)

func SetUpServer() *gin.Engine {
	r := gin.Default() // gin呼び出し

	// ---------------
	// 募集要項CRUD
	// ---------------
	r.GET("/recruitments", controller.GetRecruitments)         // 募集要項一覧を表示
	r.GET("/recruitment/:id", controller.GetRecruitmentById)   // 指定した募集要項を表示
	r.POST("/recruitment", controller.PostRecruitment)         // 募集要項を投稿
	r.PATCH("/recruitment/:id", controller.PatchRecruitment)   // 募集要項を更新
	r.DELETE("/recruitment/:id", controller.DeleteRecruitment) // 募集要項を削除

	// ---------------
	// ユーザーCRUD
	// ---------------
	r.GET("/users", controller.GetUsers)         // ユーザー一覧を表示
	r.GET("/user/:id", controller.GetUserById)   // 指定したユーザーを表示
	r.POST("/user", controller.PostUser)         // ユーザーを追加
	r.PATCH("/user/:id", controller.PatchUser)   // ユーザー情報を更新
	r.DELETE("/user/:id", controller.DeleteUser) // ユーザーを削除

	return r
}

func main() {
	log.Fatalln("Server start")
	SetUpServer().Run()
}
