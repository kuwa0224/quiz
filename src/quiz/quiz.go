package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	router := gin.Default()

	router.LoadHTMLFiles(
		"templates/top.html",
		"templates/admin/top.html",
	)
	// router.LoadHTMLGlob("templates/admin/top.html")
	router.GET("/", top)
	router.GET("/admin", admin)

	http.Handle("/", router)
}

func top(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	log.Infof(ctx, "show top")

	c.HTML(http.StatusOK, "top.html", gin.H{})
}

func admin(c *gin.Context) {

	// TODO 設定されたクイズの一覧を表示

	c.HTML(http.StatusOK, "admin/top.html", gin.H{})
}
