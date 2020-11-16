package main

import (
	"github.com/gin-gonic/gin"
	"jwt-example"
	"log"
	"net/http"
)

func main()  {
	e := gin.Default()
	e.GET("/", jwt_example.IssueToken)

	e.GET("/secret", jwt_example.JWT(), handleSecret)
	log.Fatal(e.Run(":8090"))
}

func handleSecret(c *gin.Context) {
	c.String(http.StatusOK, "secret content")
}