package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thejunghare/webdec/news"
)

func searchHandler(c *gin.Context) {
	rawURL := c.Request.URL.String()

	u, err := url.Parse(rawURL)
	if err != nil {

		return
	}

	c.String(http.StatusOK, rawURL) // <- returns the url typed
	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	fmt.Println(searchQuery)
	fmt.Println(page)
}

func main() {

	// ** A Gin application always starts by defining a router typed RouterGroup; by default,
	// ** we can get an instance of RouterGroup by calling gin.Default():
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	// ** Print hello on root dic
	r.GET("/", func(c *gin.Context) {
		// c.String(200, "Hello!")
		// Run the index.html on root
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Message": "Rendering Page",
		})
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// lastname := c.Query("lastname") // <- shortcut for c.Request.URL.Query().Get("lastname")

		// c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		c.String(http.StatusOK, "Hello %s ", firstname)
	})

	r.GET("/search", searchHandler)

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apikey must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myClient, apikey, 20)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong!")
	})

	// ** custom Template renderer
	html := template.Must(template.ParseFiles("index.html"))
	// Execute
	r.SetHTMLTemplate(html)

	// ** Run the web server
	err = r.Run()
	if err != nil {
		panic("Failed to start the gin server" + err.Error())
	}

	// r.Run(":8081") <- customized port
}
