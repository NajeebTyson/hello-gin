package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type validateStructExample struct {
	StrData     string `json:"str_data" binding:"required"`
	NumericData int    `json:"num_data" binding:"required,gt=0"`
	BooleanData bool   `json:"bool_data" binding:"required"`
	ListData    []int  `json:"list_data" binding:"required"`
}

// Validate validation example
func Validate(c *gin.Context) {
	var structExample validateStructExample
	err := c.BindJSON(&structExample)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Printf("Data: %v\n", structExample)
	c.JSON(http.StatusOK, gin.H{
		"message": structExample,
	})
}
