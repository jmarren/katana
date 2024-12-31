package data

import (
	// "github.com/jmarren/katana/src"
	"net/http"
	// "github.com/a-h/templ"
)

type AboutBodyData struct {
	Num int
}

type AboutHeadData struct{}

type AboutBodyCtr struct {
	R   *http.Request
	Num int
}

func AboutBody(c AboutBodyCtr) *AboutBodyData {
	// username := c.R.PathValue("username")
	return &AboutBodyData{
		Num: c.Num,
	}
}

type AboutHeadCtr struct {
}

func AboutHead(c AboutHeadCtr) *AboutHeadData {
	return &AboutHeadData{}
}
