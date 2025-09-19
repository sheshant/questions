package main 


import (
	"github.com/sheshant/questions/initializers"
	"github.com/sheshant/questions/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Quiz{},
		&models.Question{},
	)
}
