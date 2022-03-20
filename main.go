package main

import (
	"github.com/SteveYi-LAB/WHOIS-Search/internal/whoisSearch"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func webServer(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func whoisServer(c *gin.Context) {
	IRR := c.Param("IRR")
	target := c.Param("target")

	result := whoisSearch.IRR_DB(IRR, target)
	c.Data(200, "text/plain; charset=UTF-8", []byte(result))
}

func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

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

	router.NoRoute(pageNotAvailable)

	router.Run(":30010")
}
