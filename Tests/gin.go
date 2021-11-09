package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.Run(":8080") // listen and GO on 0.0.0.0:808
	router.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

}