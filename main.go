package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/likexian/whois-go"
)

func webServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	r.ParseForm()
	fmt.Println("\nIP Address: " + getIP(r))
	fmt.Println("method:", r.Method)
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path

	exist := "NULL"
	if _, err := os.Stat(p); err == nil {
		exist = "1"
	} else if os.IsNotExist(err) {
		exist = "0"
	}

	if p == "./whois" || p == "./whois/" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		Target := "NULL"
		for key, values := range r.Form {
			if key == "target" {
				Target = values[0]
			}
		}
		fmt.Println("Search:", Target)
		io.WriteString(w, requestWhois(Target))
	} else if r.Method == "GET" {
		if exist == "1" {
			if p == "./" {
				p = "./index.html"
			}
			http.ServeFile(w, r, p)
		} else {
			c, err := ioutil.ReadFile("./404.html")
			if err != nil {
				fmt.Println(err)
			}
			w.Write(c)
		}
	}
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func requestWhois(Target string) string {
	result, err := whois.Whois(Target)
	if err != nil {
		fmt.Println(result)
	}
	return result
}

func main() {
	http.HandleFunc("/", webServer)

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("SteveYi Whois Service\n")
	fmt.Print("Port listing at 30010\n")
	fmt.Print("Repo: https://github.com/SteveYi-LAB/SteveYi-Whois\n")
	fmt.Print("Author: SteveYi\n")
	fmt.Print("Demo: https://whois.steveyi.net\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":30010", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
