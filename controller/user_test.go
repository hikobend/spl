package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestGetUsers(t *testing.T) {
	r := SetUpRouter()                              // ルーターを作成
	r.GET("/users", GetUsers)                       // ユーザー一覧を表示
	req, _ := http.NewRequest("GET", "/users", nil) // リクエスト作成

	w := httptest.NewRecorder() // テスト用サーバー作成

	r.ServeHTTP(w, req) // テストサーバーとリクエストを受け取る

	var users []user // []userをusersで受け取る

	json.Unmarshal(w.Body.Bytes(), &users) // jsonを構造体に変化 MarshalとUmmarshal https://www.asobou.co.jp/blog/web/marshal-unmarshal

	// assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(users))
}

func TestPostUser(t *testing.T) {
	r := SetUpRouter()        // ルーター作成
	r.POST("/user", PostUser) // ユーザーを追加
	user := user{             // 追加するユーザーの設定値
		ID:       4,
		Name:     "player4",
		PlayTime: "1000000000000時間",
	}
	v, _ := json.Marshal(user)                                     // 構造体をjsonに変換 配列
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(v)) // リクエスト作成
	w := httptest.NewRecorder()                                    // テストサーバー
	r.ServeHTTP(w, req)                                            // テストサーバーとリクエストを受け取る

	assert.Equal(t, http.StatusCreated, w.Code) // assert
}
