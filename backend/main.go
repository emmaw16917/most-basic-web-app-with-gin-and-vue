package main

import (
	"net/http"

	"github.com/emmaw16917/basic-wiki-with-gin-and-vue/backend/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getHandler(c *gin.Context) {
	title := c.Param("title")
	page, err := service.LoadPage(title)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "page not found", "details": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, page)
}
func saveHandler(c *gin.Context) {
	var page service.Page
	if err := c.BindJSON(&page); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	if err := page.Save(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "message": "page saved successfully"})
}

func main() {
	router := gin.Default()

	router.Use(cors.Default()) // 允许跨域请求

	router.GET("/api/page/:title", getHandler)
	router.POST("/api/page/save", saveHandler)

	router.Run("localhost:8080")
}
