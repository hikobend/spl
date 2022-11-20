package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 募集のテーブル
type recruitment struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Need  string `json:"need"`
}

// 仮データ作成
var recruitments = []recruitment{
	{
		ID:    1,
		Title: "掲示板1",
		Need:  "募集要項1",
	},
	{
		ID:    2,
		Title: "掲示板2",
		Need:  "募集要項2",
	},
	{
		ID:    3,
		Title: "掲示板3",
		Need:  "募集要項3",
	},
}

func main() {
	r := gin.Default() // gin呼び出し

	r.GET("/recruitments", getRecruitments)         // 募集要項一覧を表示
	r.GET("/recruitment/:id", getRecruitmentById)   // 指定した募集要項を表示
	r.POST("/recruitment", postRecruitment)         // 募集要項を投稿
	r.PATCH("/recruitment/:id", patchRecruitment)   // 募集要項を更新
	r.DELETE("/recruitment/:id", deleteRecruitment) // 募集要項を削除

	r.Run("localhost:8080") // localhost:8080で実行
}

func getRecruitments(c *gin.Context) {
	c.JSON(http.StatusOK, recruitments) // 一覧を表示。JSONメソッドを使用。IndentedJSONメソッドは一旦保留。https://qiita.com/holy_engineer/items/86feea24a1b563ea37aa
}

func getRecruitmentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // idを取得する。それを数値に変換
	if err != nil {
		log.Fatalln(err) // エラーハンドリング
	}

	for _, recruitment := range recruitments { // 第二引数をrecruitmentとしている。ベストプラクティスによると短い文字(rなど)を使用するべきだが、今回は自分がわかるようにする。
		if recruitment.ID == id {
			c.JSON(http.StatusOK, recruitment) // recruitmentを募集テーブルに表示、なくても更新される
			return                             // 脱出
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "見つかりませんでした💦"}) // idがなかった場合表示

}

// https://qiita.com/ko-watanabe/items/64134c0a3871856fdc17
func postRecruitment(c *gin.Context) {
	var newRecruitment recruitment // newRecruitmentを募集テーブルにて作成

	err := c.BindJSON(&newRecruitment) // 渡した構造体ポインターをバインド
	if err != nil {
		log.Fatalln(err) // エラーハンドリング
	}

	recruitments = append(recruitments, newRecruitment) // newRecruitmentをrecruitmentsに追加
	c.JSON(http.StatusCreated, newRecruitment)          //　追加したデータを表示。なくても更新される。
}

func patchRecruitment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id")) // idを数値に変換
	if err != nil {
		log.Fatalln(err) // エラーハンドリング
	}

	var patchRecruitment recruitment // patchRecruitmentを募集テーブルにて作成
	patchRecruitment.ID = id         // patchRecruitmentのIDをidをする

	if err = c.BindJSON(&patchRecruitment); err != nil { // JSONのリクエストボディをバインド
		log.Fatalln(err) // エラーハンドリング
	}

	for i, recruitment := range recruitments {
		if recruitment.ID == id { // recruitmentのIDとidが一致していたら
			recruitments[i] = patchRecruitment      // i番目のrecruitmentsをpatchRecruitmentに更新
			c.JSON(http.StatusOK, patchRecruitment) // patchRecruitmentを募集テーブルに表示、なくても更新される
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "見つかりませんでした💦"}) // idがなかった場合表示
}

func deleteRecruitment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	for i, recruitment := range recruitments {
		if recruitment.ID == id {
			recruitments = append(recruitments[:i], recruitments[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "投稿を削除しました"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "見つかりませんでした💦"})

}
