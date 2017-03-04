package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"gopkg.in/gin-gonic/gin.v1"
)

type Answer struct {
	Description string `json:"description" binding:"required"`
	Correct     bool   `json:"correct" binding:"required"`
}

type Quiz struct {
	Question   string   `json:"question" binding:"required"`
	Answers    []Answer `json:"answers" binding:"required"`
	AnswerNote string   `json:"answerNote" binding:"required"`
}

func init() {
	router := gin.Default()

	router.LoadHTMLFiles(
		"templates/top.html",
		"templates/admin/top.html",
		"templates/admin/new.html",
	)
	// router.LoadHTMLGlob("templates/admin/top.html")
	router.GET("/", Top)
	router.GET("/admin", QuizList)
	router.GET("/admin/new", QuizNewForm)
	router.POST("/admin/new", QuizNew)

	http.Handle("/", router)
}

func Top(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	log.Infof(ctx, "show top")

	c.HTML(http.StatusOK, "top.html", gin.H{})
}

func QuizList(c *gin.Context) {

	// TODO 設定されたクイズの一覧を表示

	c.HTML(http.StatusOK, "admin/top.html", gin.H{})
}

func QuizNewForm(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/new.html", gin.H{})
}

func QuizNew(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)

	var quiz Quiz
	if err := c.BindJSON(&quiz); err != nil {
		log.Errorf(ctx, "Failed to bind : %v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
