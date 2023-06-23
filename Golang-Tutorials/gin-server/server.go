package main

import "github.com/gin-gonic/gin"

func main() {
	localhost := gin.Default() // default PORT=8080

	// A handler for GET request on /
	localhost.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	// Request at localhost:8080/users
	localhost.GET("/users", func(c *gin.Context) {
		users := []string{"Alice", "BOB", "Chalie"}
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	// Request at localhost:8080/prasad
	localhost.GET("/prasad", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Prasad",
		})
	})

	// POST request
	localhost.POST("/users", func(c *gin.Context) {
		var user struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(201, gin.H{
			"Message": "User created",
		})
	})

	localhost.Run() // listen and serve
}
