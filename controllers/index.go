package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Nodes": router.nodeList.Nodes,
	})
}
