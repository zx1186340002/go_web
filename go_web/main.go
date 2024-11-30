package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"addrss"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", GetPerson)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetPerson(c *gin.Context) {
	person := Person{
		Name:    "舒镐",
		Age:     18,
		Address: "背景",
	}
	fmt.Println(person)
	c.JSON(http.StatusOK, person)
}
