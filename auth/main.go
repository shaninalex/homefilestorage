package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"
)

var (
	COOCKIE_NAME = os.Getenv("AUTH_SERVER_SESSION_NAME")
)

func main() {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: os.Getenv("AUTH_SERVER"), // Kratos Public API
		},
	}
	apiClient := client.NewAPIClient(configuration)

	router := gin.Default()
	router.GET("/identity", func(c *gin.Context) {
		cookie, err := c.Cookie(COOCKIE_NAME)
		if err != nil {
			log.Println(err)
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		resp, r, err := apiClient.FrontendApi.ToSession(context.Background()).Cookie(fmt.Sprintf("%s=%s", COOCKIE_NAME, cookie)).Execute()
		if err != nil {
			log.Printf("Error when calling `FrontendApi.ToSession``: %v\n", err)
			log.Printf("Full HTTP response: %v\n", r)
			c.JSON(400, gin.H{
				"error": "Server error",
			})
			return
		}

		// response
		// {
		// 	"email": "test@test.com",
		// 	"name": {
		// 		"first": "Firstname",
		// 		"last": "Lastname"
		// 	}
		// }
		c.JSON(200, resp.Identity.Traits)
	})

	portInt, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	router.Run(fmt.Sprintf(":%d", portInt))
}
