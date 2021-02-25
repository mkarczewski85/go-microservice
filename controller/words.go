package controller

import (
	"net/http"
	"words-microservice/model"
	"words-microservice/service"

	"github.com/gin-gonic/gin"
)

// POST
func SearchForWordsHandler(c *gin.Context) {
	var input model.WordsSearchInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := service.FindMatchedWords(&input)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
