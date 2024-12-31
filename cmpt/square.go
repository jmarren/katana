package cmpt

import (
	"github.com/jmarren/katana/data"
	// "github.com/jmarren/katana/render"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
)

func Square() *src.Component {
	squareBodyCtr := src.EltCtr[struct{}, *data.SquareBodyData]{
		Ctr:       data.SquareBody,
		TemplFunc: templates.SquareBody,
	}

	squareHeadCtr := src.EltCtr[struct{}, *data.SquareHeadData]{
		Ctr:       data.SquareHead,
		TemplFunc: templates.SquareHead,
	}

	squareFunc := src.NewComponent(squareBodyCtr, squareHeadCtr)
	square := squareFunc(data.EmptyData{}, data.EmptyData{})
	return square
}
