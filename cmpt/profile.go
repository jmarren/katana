package cmpt

import (
	"github.com/jmarren/katana/data"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
	"net/http"
)

func Profile(r *http.Request, square *src.Component, circle *src.Component) *src.Component {
	profileHeadCtr := src.EltCtr[data.ProfileHeadCtr, *data.ProfileHeadData]{
		Ctr:       data.ProfileHead,
		TemplFunc: templates.ProfileHead,
	}

	profileBodyCtr := src.EltCtr[data.ProfileBodyCtr, *data.ProfileBodyData]{
		Ctr:       data.ProfileBody,
		TemplFunc: templates.ProfileBody,
	}

	profileFunc := src.NewComponent("profile", profileBodyCtr, profileHeadCtr)

	profile := profileFunc(
		data.ProfileBodyCtr{
			R:      r,
			Square: square.Body,
			Circle: circle.Body,
		},
		data.ProfileHeadCtr{
			Square: square.Head,
			Circle: circle.Head,
		},
	)
	return profile
}
