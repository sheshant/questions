package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sheshant/questions/initializers"
	"github.com/sheshant/questions/models"
)

var body struct {
	Title 	string
}

func QuizCreate(c * gin.Context) {
	c.Bind(&body)

	todo := models.Quiz{Title: body.Title}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
  		return
	}
	c.JSON(200, todo)

}

func QuizGet(c * gin.Context) {
	id := c.Param("id")
	var quiz models.Quiz
	initializers.DB.First(&quiz, id) 
	c.JSON(200, quiz)
}

func QuizGetAll(c * gin.Context) {
	var quizes []models.Quiz
	initializers.DB.Find(&quizes)
	c.JSON(200, gin.H{
  		"quizes": quizes,
	})

	
}

func QuizUpdate(c * gin.Context) {
	return
}

func QuizDelete(c * gin.Context) {
	return
}
