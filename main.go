package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/likexian/whois"
)

func webServer(ctx *gin.Context) {
	ctx.HTML(200, "index.tmpl", nil)
}

func whoisServer(ctx *gin.Context) {
	Target := ctx.Param("target")
	result, err := whois.Whois(Target)
	if err != nil {
		fmt.Println(result)
	}
	ctx.String(200, result)
}

func whoisRADB(ctx *gin.Context) {
	Target := ctx.Param("target")
	result, err := whois.Whois(Target, "whois.radb.net")
	if err != nil {
		fmt.Println(result)
	}
	ctx.String(200, result)
}

func whoisPOST(ctx *gin.Context) {
	Target := ctx.PostForm("target")
	result, err := whois.Whois(Target)
	if err != nil {
		fmt.Println(result)
	}
	ctx.String(200, result)
}

func whoisRADBPOST(ctx *gin.Context) {
	Target := ctx.PostForm("target")
	result, err := whois.Whois(Target, "whois.radb.net")
	if err != nil {
		fmt.Println(result)
	}
	ctx.String(200, result)
}

func pageNotAvailable(ctx *gin.Context) {
	ctx.HTML(404, "404.tmpl", nil)
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
	router.GET("/whois/:target", whoisServer)
	router.POST("/whois/", whoisPOST)
	router.GET("/RADB/:target", whoisRADB)
	router.POST("/RADB/", whoisRADBPOST)

	router.NoRoute(pageNotAvailable)

	router.Run(":30010")
}
