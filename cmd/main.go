package main

import (
	"fmt"

	"github.com/jmarren/katana/build"
	"net/http"
)

func main() {
	fmt.Println("running")
	// about := components.Run()
	node, err := build.GetNodeFactory("index")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	http.HandleFunc("/{page}", node.Render)
	http.ListenAndServe(":8080", nil)
}
