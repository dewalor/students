package main

import (
	"github.com/gin-gonic/gin"

	"students/models"
	"students/controllers"

)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/students", controllers.FindStudents)
    r.POST("/students", controllers.CreateStudent)
	
	r.Run()
}
