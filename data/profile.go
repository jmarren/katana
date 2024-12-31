package data

import (
	// "github.com/jmarren/katana/src"
	"net/http"

	"github.com/a-h/templ"
)

type ProfileBodyData struct {
	Username string
	Circle   *templ.Component
	Square   *templ.Component
}

type ProfileHeadData struct {
	Circle *templ.Component
	Square *templ.Component
}

type ProfileBodyCtr struct {
	R      *http.Request
	Circle *templ.Component
	Square *templ.Component
}

func ProfileBody(c ProfileBodyCtr) *ProfileBodyData {
	username := c.R.PathValue("username")
	return &ProfileBodyData{
		Username: username,
		Circle:   c.Circle,
		Square:   c.Square,
	}
}

type ProfileHeadCtr struct {
	Circle *templ.Component
	Square *templ.Component
}

func ProfileHead(c ProfileHeadCtr) *ProfileHeadData {
	return &ProfileHeadData{
		Circle: c.Circle,
		Square: c.Square,
	}
}
