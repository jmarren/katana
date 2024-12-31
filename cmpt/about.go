package cmpt

import (
	"github.com/jmarren/katana/data"
	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
	"net/http"
)

func About(r *http.Request, num int) *src.Component {
	aboutHeadCtr := src.EltCtr[data.AboutHeadCtr, *data.AboutHeadData]{
		Ctr:       data.AboutHead,
		TemplFunc: templates.AboutHead,
	}

	aboutBodyCtr := src.EltCtr[data.AboutBodyCtr, *data.AboutBodyData]{
		Ctr:       data.AboutBody,
		TemplFunc: templates.AboutBody,
	}

	aboutFunc := src.NewComponent("about", aboutBodyCtr, aboutHeadCtr)

	about := aboutFunc(
		data.AboutBodyCtr{
			R:   r,
			Num: num,
		},
		data.AboutHeadCtr{},
	)
	return about
}
