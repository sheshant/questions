package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sheshant/questions/initializers"
	"github.com/sheshant/questions/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.QuestionsRoutes(r)
	r.Run()
}