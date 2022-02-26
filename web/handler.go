package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func GetContentById(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func PostContentById(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func PatchContentById(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func DeleteContentById(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}
