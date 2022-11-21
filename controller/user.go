package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PlayTime string `json:"playTime"`
}

var users = []user{
	{
		ID:       1,
		Name:     "player1",
		PlayTime: "100時間",
	},
	{
		ID:       2,
		Name:     "player2",
		PlayTime: "10000時間",
	},
	{
		ID:       3,
		Name:     "player3",
		PlayTime: "10時間",
	},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "見つかりませんでした💦"})
}

func PostUser(c *gin.Context) {
	var newUser user

	err := c.BindJSON(&newUser)
	if err != nil {
		log.Fatalln(err)
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func PatchUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	var patchUser user
	patchUser.ID = id

	if err = c.BindJSON(&patchUser); err != nil {
		log.Fatalln(err)
	}

	for i, user := range users {
		if user.ID == id {
			users[i] = patchUser
			c.JSON(http.StatusOK, patchUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "見つかりませんでした💦"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "ユーザーを削除しました"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "見つかりませんでした💦"})
}
