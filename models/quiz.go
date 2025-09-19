package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID 			int 	`gorm:"unique;primaryKey;autoIncrement"`
	Title 		string
	Questions   []Question `gorm:"foreignKey:QuizID;references:ID"`
}

