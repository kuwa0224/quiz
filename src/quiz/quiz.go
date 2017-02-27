package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", top)

	http.Handle("/", router)
}

func top(c *gin.Context) {
	ctx := appengine.WithContext(c, c.Request)
	log.Infof(ctx, "OK")

	c.HTML(http.StatusOK, "top.html", gin.H{
		"name": "Takeshi",
	})
}
