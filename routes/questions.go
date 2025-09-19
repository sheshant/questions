package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sheshant/questions/controllers"
)

func QuestionsRoutes(r *gin.Engine) {
	QuestionsGroup := r.Group("/questions")
	{
		QuestionsGroup.POST("", controllers.QuestionCreate)       // Create
		QuestionsGroup.GET("", controllers.QuestionGetAll)         // Read all
		QuestionsGroup.GET("/:id", controllers.QuestionGet)      // Read One
		QuestionsGroup.PUT("/:id", controllers.QuestionUpdate)    // Update
		QuestionsGroup.DELETE("/:id", controllers.QuestionDelete) // Delete
	}
	QuizGroup := r.Group("/quiz")
	{
		QuizGroup.POST("", controllers.QuizCreate)       // Create
		QuizGroup.GET("", controllers.QuizGetAll)         // Read all
		QuizGroup.GET("/:id", controllers.QuizGet)      // Read One
		QuizGroup.PUT("/:id", controllers.QuizUpdate)    // Update
		QuizGroup.DELETE("/:id", controllers.QuizDelete) // Delete
	}
}
