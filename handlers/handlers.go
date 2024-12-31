package handlers

import (
	"github.com/jmarren/katana/cmpt"
	"github.com/jmarren/katana/data"
	// "github.com/jmarren/katana/render"
	"net/http"

	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	profileHeadCtr := src.EltCtr[data.ProfileHeadCtr, *data.ProfileHeadData]{
		Ctr:       data.ProfileHead,
		TemplFunc: templates.ProfileHead,
	}

	profileBodyCtr := src.EltCtr[data.ProfileBodyCtr, *data.ProfileBodyData]{
		Ctr:       data.ProfileBody,
		TemplFunc: templates.ProfileBody,
	}

	profileFunc := src.NewComponent(profileBodyCtr, profileHeadCtr)

	square := cmpt.Square()

	profile := profileFunc(
		data.ProfileBodyCtr{
			R:     r,
			Child: square.Body,
		},
		data.ProfileHeadCtr{
			Child: square.Head,
		},
	)

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
