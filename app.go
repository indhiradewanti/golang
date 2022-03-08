package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	service := gin.Default()
	service.POST("/name", PostName)
	service.GET("/name", GetName)
	service.Run(":8000")
}

func GetName(c *gin.Context) {
	name := "Indhira"
	c.IndentedJSON(http.StatusOK, gin.H{"name": name})
}

func PostKeyword(keyword string) (response string) {
	fmt.Println("MASUK POST KEYWORD", keyword)
	switch strings.ToLower(keyword) {
	case "indhira":
		return "dhira"
	case "ramzy":
		return "zy"
	default:
		return "unknown"
	}
}

type Key struct {
	Keyword string `json:"keyword"`
}

func PostName(c *gin.Context) {
	input := Key{}

	// fmt.Println("MASUK POST NAME", c.BindJSON(&input))

	if err := c.BindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	var output string
	switch strings.ToLower(input.Keyword) {
	case "indhira":
		output = "dhira"
	case "ramzy":
		output = "ramzy"
	default:
		output = "unknown"
	}
	// output := PostKeyword(c.PostForm("keyword"))

	c.IndentedJSON(http.StatusOK, gin.H{"name" : output})
}
