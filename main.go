package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hikobend/spl/controller"
)

func main() {
	r := gin.Default() // gin呼び出し

	r.GET("/recruitments", controller.GetRecruitments)         // 募集要項一覧を表示
	r.GET("/recruitment/:id", controller.GetRecruitmentById)   // 指定した募集要項を表示
	r.POST("/recruitment", controller.PostRecruitment)         // 募集要項を投稿
	r.PATCH("/recruitment/:id", controller.PatchRecruitment)   // 募集要項を更新
	r.DELETE("/recruitment/:id", controller.DeleteRecruitment) // 募集要項を削除

	r.Run("localhost:8080") // localhost:8080で実行
}
