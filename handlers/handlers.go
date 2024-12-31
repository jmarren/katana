package handlers

import (
	"github.com/jmarren/katana/data"
	// "github.com/jmarren/katana/render"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// profile := src.Wrap()
	square := src.NewComponent(data.Square, templates.SquareBody, templates.SquareHead)(data.EmptyData{})
	profileWSquare := square.Wrap(templates.ProfileHead, templates.ProfileBody)
	profileWSquare.Render(w, r, templates.Basefunc)

	// profile := src.NewComponent(data.Profile, templates.ProfileBody, templates.ProfileHead)(r)
	// profile.Render(w, r, templates.Basefunc)
}

func ProfileWithSquareHandler(w http.ResponseWriter, r *http.Request) {

}

// profile :=
// profileWSquare := src.NewComponent()
/*
	WANT: profileWithSquare
		-- Profile src.Component
		-- with data

*/

// }
