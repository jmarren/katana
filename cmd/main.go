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
	// about := components.Run()
	// node, err := build.GetNodeFactory("index")
	// if err != nil {
	// 	fmt.Printf("error: %s\n", err)
	// }

	http.HandleFunc("/{username}", handlers.ProfileHandler)
	http.ListenAndServe(":8080", nil)
}
