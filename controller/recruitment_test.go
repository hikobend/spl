package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecruitments(t *testing.T) {
	r := SetUpRouter()                      // ルーターを作成
	r.GET("/recruitments", GetRecruitments) // 募集要項一覧を表示

	req, _ := http.NewRequest("GET", "/recruitments", nil) // リクエスト作成

	w := httptest.NewRecorder() // テスト用サーバー作成

	r.ServeHTTP(w, req) // テストサーバーとリクエストを受け取る

	var recruitments []recruitment // []recruitmentをrecruitmentsで受け取る

	json.Unmarshal(w.Body.Bytes(), &recruitments) // jsonを構造体に変化 MarshalとUmmarshal https://www.asobou.co.jp/blog/web/marshal-unmarshal

	// assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(recruitments))
}

func TestPostRecruitment(t *testing.T) {
	r := SetUpRouter()                      // ルーター作成
	r.POST("/recruitment", PostRecruitment) // 募集要項を投稿
	recruitment := recruitment{             // 追加するユーザーの設定値
		ID:    4,
		Title: "掲示板4",
		Need:  "募集要項4",
	}
	v, _ := json.Marshal(recruitment)                                     // 構造体をjsonに変換 配列
	req, _ := http.NewRequest("POST", "/recruitment", bytes.NewBuffer(v)) // リクエスト作成
	w := httptest.NewRecorder()                                           // テストサーバー
	r.ServeHTTP(w, req)                                                   // テストサーバーとリクエストを受け取る

	assert.Equal(t, http.StatusCreated, w.Code) // assert
}
