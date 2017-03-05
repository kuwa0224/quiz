package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"gopkg.in/gin-gonic/gin.v1"
)

type Answer struct {
	Description string `json:"description" binding:"required"`
	Correct     bool   `json:"correct" binding:"required"`
}

type Quiz struct {
	ID         *datastore.Key `json:"id"`
	Question   string         `json:"question" binding:"required"`
	Answers    []Answer       `json:"answers" binding:"required"`
	AnswerNote string         `json:"answerNote" binding:"required"`
}

type QuizForm struct {
	Quiz      Quiz `json:"quiz" binding:"required"`
	CheckOnly bool `json:"checkOnly" binding:"exists"`
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
	ctx := appengine.NewContext(c.Request)

	// TODO 設定されたクイズの一覧を表示
	q := datastore.NewQuery("Quiz")

	var quizList []Quiz
	q.GetAll(ctx, &quizList)

	log.Infof(ctx, "quiz List: %v", quizList)

	c.HTML(http.StatusOK, "admin/top.html", gin.H{"QuizList": quizList})
}

func QuizNewForm(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/new.html", gin.H{})
}

func QuizNew(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)

	var qf QuizForm
	if err := c.BindJSON(&qf); err != nil {
		log.Errorf(ctx, "Failed to bind : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if qf.CheckOnly {
		log.Infof(ctx, "Request is only for check")
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	log.Infof(ctx, "Post actually")
	key := datastore.NewIncompleteKey(ctx, "Quiz", nil)
	qf.Quiz.ID = key
	key, err := datastore.Put(ctx, key, &qf.Quiz)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
