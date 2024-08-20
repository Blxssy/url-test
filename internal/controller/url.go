package controller

import (
	"github.com/Blxssy/url-test/internal/container"
	"github.com/Blxssy/url-test/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type URLController interface {
	CreateShortURL(c *gin.Context)
	Redirect(c *gin.Context)
}

type urlController struct {
	container container.Container
	service   service.URLService
}

func NewURLController(container container.Container) URLController {
	return &urlController{
		container: container,
		service:   service.NewURLService(container),
	}
}

func (uc *urlController) CreateShortURL(c *gin.Context) {
	var input struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cant get url"})
		return
	}

	shortURL, err := uc.service.SaveURL(input.URL, uc.container.GetConfig().URLSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cant create url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortURL})
}

func (uc *urlController) Redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")

	originalURL, err := uc.service.GetOriginalURL(shortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cant get original url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"original_url": originalURL})
}
