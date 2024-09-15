package routes

import (
	"fastfishback/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterLessonRoutes(router *gin.Engine) {
	lessonRoutes := router.Group("/lessons")
	{
		lessonRoutes.GET("/", controllers.GetLessons)
		lessonRoutes.GET("/:id", controllers.GetLessonByID)
		lessonRoutes.POST("/", controllers.CreateLesson)
		lessonRoutes.PUT("/:id", controllers.UpdateLesson)
		lessonRoutes.DELETE("/:id", controllers.DeleteLesson)
	}
}
