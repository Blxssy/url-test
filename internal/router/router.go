package router

import (
	"github.com/Blxssy/url-test/internal/container"
	"github.com/Blxssy/url-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.Engine, container container.Container) {
	setURLController(g, container)
}

func setURLController(g *gin.Engine, container container.Container) {
	url := controller.NewURLController(container)

	g.POST("/", url.CreateShortURL)
	g.GET("/:shortURL", url.Redirect)
}
