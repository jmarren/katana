package cmpt

import (
	"github.com/jmarren/katana/data"
	// "github.com/jmarren/katana/render"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
)

func Square(color string) *src.Component {
	squareBodyCtr := src.EltCtr[data.SquareBodyData, *data.SquareBodyData]{
		Ctr:       data.SquareBody,
		TemplFunc: templates.SquareBody,
	}

	squareHeadCtr := src.EltCtr[data.SquareHeadData, *data.SquareHeadData]{
		Ctr:       data.SquareHead,
		TemplFunc: templates.SquareHead,
	}

	squareFunc := src.NewComponent("square", squareBodyCtr, squareHeadCtr)
	square := squareFunc(data.SquareBodyData{
		Color: color,
	}, data.SquareHeadData{})
	return square
}
