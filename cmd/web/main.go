package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/catehulu/factorio-calculator/internal/models"
)

func main() {
	var path string
	var port string
	flag.StringVar(&path, "path", "/", "Provide project path as an absolute path")
	flag.StringVar(&port, "port", ":8080", "Provide project port")
	flag.Parse()
	models.InitItem(path)
	models.InitRecipe(path)

	server := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	fmt.Printf("Started server on port %v with path %v", port, path)
	err := server.ListenAndServe()
	log.Fatal(err)
}
