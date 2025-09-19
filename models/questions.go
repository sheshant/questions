package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID 			int 	`gorm:"unique;primaryKey;autoIncrement"`
	Text 		string
	Option_a 	string
	Option_b 	string
	Option_c 	string
	Option_d 	string
	Answer	 	string
	QuizID    	int    	`gorm:"index;nil=True"`
}

