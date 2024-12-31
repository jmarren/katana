package main

import (
	"fmt"

	"net/http"

	"github.com/jmarren/katana/handlers"
	// "github.com/jmarren/katana/data"
	// "github.com/jmarren/katana/src"
	// "github.com/jmarren/katana/templates"
)

func main() {
	fmt.Println("running")

	http.HandleFunc("GET /{page}", handlers.PageHandler)
	http.HandleFunc("GET /square/{color}", handlers.SquareHandler)
	http.HandleFunc("GET /circle/{color}", handlers.CircleHandler)

	http.ListenAndServe(":8080", nil)
}
