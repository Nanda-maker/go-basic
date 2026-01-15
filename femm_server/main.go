package main

import (
	"fmt"
	"go/femm_server/api"
	"go/femm_server/data"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func handleHello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from a Go Program!!!"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	html.Execute(w, data.GetAll())

}

func main() {

	go func() {
	router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run() // listens on 0.0.0.0:8080 by default
}()


	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello) //pass func as an argument in Go
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)

	if err == nil {
		fmt.Println("Error while opening the server")
	}

}
