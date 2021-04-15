package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/likexian/whois"
)

func webServer(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func whoisServer(c *gin.Context) {
	Target := strings.Replace(c.Param("target"), "/", "", 1)
	result, err := whois.Whois(Target)
	if err != nil {
		fmt.Println(result)
	}
	c.Data(200, "text/plain; charset=UTF-8", []byte(result))
}

func whoisRADB(c *gin.Context) {
	Target := strings.Replace(c.Param("target"), "/", "", 1)
	result, err := whois.Whois(Target, "whois.radb.net")
	if err != nil {
		fmt.Println(result)
	}
	c.Data(200, "text/plain; charset=UTF-8", []byte(result))
}

func whoisPOST(c *gin.Context) {
	Target := strings.Replace(c.Param("target"), "/", "", 1)
	result, err := whois.Whois(Target)
	if err != nil {
		fmt.Println(result)
	}
	c.Data(200, "text/plain; charset=UTF-8", []byte(result))
}

func whoisRADBPOST(c *gin.Context) {
	Target := strings.Replace(c.Param("target"), "/", "", 1)
	result, err := whois.Whois(Target, "whois.radb.net")
	if err != nil {
		fmt.Println(result)
	}
	c.Data(200, "text/plain; charset=UTF-8", []byte(result))
}

func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func main() {

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("SteveYi Whois Service\n")
	fmt.Print("Port listing at 30010\n")
	fmt.Print("Repo: https://github.com/SteveYi-LAB/SteveYi-Whois\n")
	fmt.Print("Author: SteveYi\n")
	fmt.Print("Demo: https://whois.steveyi.net\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("static/*")

	router.GET("/", webServer)
	router.GET("/whois/*target", whoisServer)
	router.POST("/whois/", whoisPOST)
	router.GET("/RADB/*target", whoisRADB)
	router.POST("/RADB/", whoisRADBPOST)

	router.NoRoute(pageNotAvailable)

	router.Run(":30010")
}
