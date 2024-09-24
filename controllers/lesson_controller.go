package controllers

import (
	"fastfishback/models"
	"fastfishback/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLessons(c *gin.Context) {
	lessons, err := services.GetAllLessons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

// GetLessonByID handles the GET request to retrieve a lesson by ID
func GetLessonByID(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	lesson, err := services.GetLessonByID(uint(lessonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// CreateLesson handles the POST request to create a new lesson
func CreateLesson(c *gin.Context) {
	var lesson models.Lesson

	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLesson, err := services.CreateLesson(&lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdLesson)
}

// UpdateLesson handles the PUT request to update an existing lesson
func UpdateLesson(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedLesson models.Lesson

	if err := c.ShouldBindJSON(&updatedLesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson, err := services.UpdateLesson(uint(lessonID), updatedLesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// DeleteLesson handles the DELETE request to delete a lesson
func DeleteLesson(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = services.DeleteLesson(uint(lessonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}
