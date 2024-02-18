package controllers

import (
	"chuckEspaNorris/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type fetchJokesResponse struct {
	Jokes []string `json:"jokes"`
}

func handlefetchJokes(context *gin.Context) {
	fmt.Println("=========================")
	jokes := services.GetRandomJokes()
	response := &fetchJokesResponse{ Jokes: jokes }
	
	// No need to parse the response, gin does it
	// jsonResponse, err := json.Marshal(response)
	// fmt.Println(string(jsonResponse))
	// if err != nil {
	// 	fmt.Println(err)
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }


	context.JSON(http.StatusOK, response)
}