package handlers

import (
	"github.com/jmarren/katana/cmpt"
	"github.com/jmarren/katana/render"
	// "github.com/jmarren/katana/templates"
	"fmt"
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	pageRequested := r.PathValue("page")
	fmt.Printf("page requested: %s\n", pageRequested)
	switch pageRequested {
	case "profile":
		ProfileHandler(w, r)
	case "about":
		AboutHandler(w, r)
	default:
		ProfileHandler(w, r)
	}

}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	squareColor := r.URL.Query().Get("square")
	if squareColor == "" {
		squareColor = "red"
	}
	profile := cmpt.Profile(r, cmpt.Square(squareColor))
	render.RenderPath(profile, w, r, "page")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	about := cmpt.About(r, 13)
	render.RenderPath(about, w, r, "page")
}

func SquareHandler(w http.ResponseWriter, r *http.Request) {
	color := r.PathValue("color")
	square := cmpt.Square(color)
	render.RenderQuery(square, w, r, color)
}
