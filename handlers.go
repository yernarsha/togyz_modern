package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"togyz_modern/models"
)

func getGames(c *gin.Context) {

	games, err := models.GetGames(10)

	checkErr(err)

	if games == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Games Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": games})
	}
}

func getGameById(c *gin.Context) {

	id := c.Param("id")
	game, err := models.GetGameById(id)

	checkErr(err)
	// if the name is blank we can assume nothing is found
	if game.WhiteName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Games Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": game})
	}
}

func options(c *gin.Context) {

	ourOptions := "HTTP/1.1 200 OK\n" +
		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Origin: http://locahost:8080\n" +
		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Headers: Content-Type\n"

	c.String(200, ourOptions)
}
