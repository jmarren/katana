package cmpt

import (
	"github.com/jmarren/katana/data"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
)

func Circle(color string) *src.Component {
	circleBodyCtr := src.EltCtr[data.CircleBodyData, *data.CircleBodyData]{
		Ctr:       data.CircleBody,
		TemplFunc: templates.CircleBody,
	}

	circleHeadCtr := src.EltCtr[data.CircleHeadData, *data.CircleHeadData]{
		Ctr:       data.CircleHead,
		TemplFunc: templates.CircleHead,
	}

	circleFunc := src.NewComponent("circle", circleBodyCtr, circleHeadCtr)
	circle := circleFunc(data.CircleBodyData{
		Color: color,
	}, data.CircleHeadData{})
	return circle
}
