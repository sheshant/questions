package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sheshant/questions/initializers"
	"github.com/sheshant/questions/models"
)

type QuestionInput struct {
        Text     string `json:"text" binding:"required"`
        OptionA  string `json:"option_a" binding:"required"`
        OptionB  string `json:"option_b" binding:"required"`
        OptionC  string `json:"option_c" binding:"required"`
        OptionD  string `json:"option_d" binding:"required"`
        Answer   string `json:"answer" binding:"required"`
        QuizID   *int   `json:"quiz_id"` // Nullable QuizID
    }

func QuestionCreate(c * gin.Context) {
    var body QuestionInput

	c.Bind(&body)

    question := models.Question{
        Text:     body.Text,
        Option_a: body.OptionA,
        Option_b: body.OptionB,
        Option_c: body.OptionC,
        Option_d: body.OptionD,
        Answer:   body.Answer,
    }
    if body.QuizID != nil {
        question.QuizID = *body.QuizID
    } // Else, QuizID remains 0, which GORM maps to NULL

    // Save to database
    if err := initializers.DB.Create(&question).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create question: " + err.Error()})
        return
    }

    c.JSON(201, gin.H{"message": "Question created successfully", "question": question})

}

func QuestionGet(c * gin.Context) {
	id := c.Param("id")
	var question models.Question
	initializers.DB.First(&question, id) 
	c.JSON(200, question)
}

func QuestionGetAll(c * gin.Context) {
	var question []models.Question
	initializers.DB.Find(&question)
	c.JSON(200, gin.H{
  		"question": question,
	})
}

func QuestionUpdate(c * gin.Context) {
	var body QuestionInput
	id := c.Param("id")
	c.Bind(&body)
	var question models.Question
	initializers.DB.First(&question, id)

	new_question := models.Question{
			Text:     body.Text,
			Option_a: body.OptionA,
			Option_b: body.OptionB,
			Option_c: body.OptionC,
			Option_d: body.OptionD,
			Answer:   body.Answer,
		}

	initializers.DB.Model(&question).Updates(new_question)
}

func QuestionDelete(c * gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Question{}, id)
	c.JSON(202, gin.H{
		"message": "Question removed Successfully",
	})
}
