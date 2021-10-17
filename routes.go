package main

import (
	"log"
	"net/http"

	"github.com/NajeebTyson/hello-gin/controller"
	"github.com/gin-gonic/gin"
)

// GetRoutes make all the routes and export them
func GetRoutes(router *gin.Engine) {

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	userRoutes := router.Group("/user")
	userRoutes.POST("/", controller.AddUser)
	userRoutes.GET("/:id", controller.GetUser)
	userRoutes.PUT("/:id", controller.UpdateUser)
	userRoutes.DELETE("/:id", controller.DeleteUser)

	wowRoutes := router.Group("/wow")
	wowRoutes.POST("/validate", controller.Validate)

	router.POST("/upload", uploadFile)
}

func uploadFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := form.File["file"]
	for _, file := range files {
		log.Println("Got file:", file.Filename)
		err := ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			log.Fatal("Unable to save file", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to save file",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Files uploaded",
	})
}
