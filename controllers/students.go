package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"students/models"
)

type CreateStudentInput struct {
	Name  string `json:"name" binding:"required"`
//	Grade int `json:"grade" binding:"required"`
	Grade int `json:"grade,string,required"`
}

// GET /students
// Get all students
func FindStudents(c *gin.Context) {
	var students []models.Student
	models.DB.Find(&students)

	c.JSON(http.StatusOK, gin.H{"data": students})
}

// POST /students
// Create new student
func CreateStudent(c *gin.Context) {
	// Validate input
	var input CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create student
	student := models.Student{Name: input.Name, Grade: input.Grade}
	models.DB.Create(&student)

	c.JSON(http.StatusOK, gin.H{"data": student})
}
