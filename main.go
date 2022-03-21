package main

import (
	"github.com/SteveYi-LAB/WHOIS-Search/internal/punycode"
	"github.com/SteveYi-LAB/WHOIS-Search/internal/tools"
	"github.com/SteveYi-LAB/WHOIS-Search/internal/whoisSearch"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Index Page
func webServer(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

// 404 Page
func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

// Query whois
func whoisServer(c *gin.Context) {
	IRR := c.Param("IRR")
	target := punycode.ConvertertoASCII(c.Param("target"))

	result := whoisSearch.IRR_DB(IRR, target)
	c.Data(200, "text/plain; charset=UTF-8", []byte(result))
}

// Query whois and return as RESTful API
func whoisServerAPI(c *gin.Context) {

	// Get String Parameter
	IRR := c.Param("IRR")
	target := punycode.ConvertertoASCII(c.Param("target"))

	result := whoisSearch.IRR_DB(IRR, target)

	// Define struct
	type whoisResult struct {
		IRR    string `json:"IRR"`
		Type   string `json:"Type"`
		Result string `json:"Result"`
	}
	var return_result whoisResult
	return_result = whoisResult{IRR, tools.CheckType(target), result}

	c.JSON(200, return_result)
}

// Gin Engine
func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.LoadHTMLGlob("static/*")

	router.GET("/", webServer)
	router.GET("/whois/:target", whoisServer)
	router.GET("/whois/:target/:IRR", whoisServer)

	router.GET("/api/whois/:target", whoisServerAPI)
	router.GET("/api/whois/:target/:IRR", whoisServerAPI)

	router.NoRoute(pageNotAvailable)

	router.Run(":30010")
}
