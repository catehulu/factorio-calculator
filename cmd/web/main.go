package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/catehulu/factorio-calculator/internal/models"
)

const portNumber string = ":8080"

func main() {
	var path string
	flag.StringVar(&path, "path", "/", "Provide project path as an absolute path")
	flag.Parse()
	models.InitItem(path)
	models.InitRecipe(path)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
