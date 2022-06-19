package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"togyz_modern/models"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("game", getGames)
		v1.GET("game/:id", getGameById)
		v1.OPTIONS("game", options)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
