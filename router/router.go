package router

import (
	"github.com/eduardylopes/link-shortener/internal/link"
	"github.com/gin-gonic/gin"
)

var req *gin.Engine

func InitRouter(linkHandler *link.Handler) {
	req = gin.Default()

	req.POST("/", linkHandler.CreateLink)
	req.GET("/:code", linkHandler.GetLinkByCode)
}

func Start() error {
	return req.Run()
}
