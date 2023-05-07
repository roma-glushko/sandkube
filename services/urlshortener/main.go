package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type LinkInfo struct {
	Url string `uri:"url" binding:"required"`
}

func main() {
	storage := map[string]string{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var link LinkInfo

		if err := c.ShouldBind(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		linkId := uuid.NewString()

		storage[linkId] = link.Url

		c.JSON(http.StatusOK, gin.H{"uuid": linkId})
	})

	router.DELETE("/:linkId", func(c *gin.Context) {
		linkId := c.Param("linkId")

		delete(storage, linkId)
	})

	router.GET("/:linkId", func(c *gin.Context) {
		linkId := c.Param("linkId")

		url, found := storage[linkId]

		if !found {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		c.Redirect(http.StatusFound, url)
	})

	router.GET("/health/", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	})

	router.Run(":8000")
}
