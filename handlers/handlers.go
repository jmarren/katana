package handlers

import (
	"github.com/jmarren/katana/data"
	// "github.com/jmarren/katana/render"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	profileFunc := src.NewComponent(data.ProfileBody, data.ProfileHead, templates.ProfileBody, templates.ProfileHead)
	squareBody := templates.SquareBody(data.SquareBody(data.EmptyData{}))
	squareHead := templates.SquareHead(data.SquareHead(data.EmptyData{}))
	profile := profileFunc(
		data.ProfileBodyCtr{
			R:     r,
			Child: &squareBody,
		},
		data.ProfileHeadCtr{
			Child: &squareHead,
		},
	)

	// profileWSquare := square.Wrap(templates.ProfileHead, templates.ProfileBody)
	profile.Render(w, r, templates.Basefunc)
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
