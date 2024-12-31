package data

import (
	// "github.com/jmarren/katana/src"
	"net/http"

	"github.com/a-h/templ"
)

type ProfileBodyData struct {
	Username string
	Child    *templ.Component
}

type ProfileHeadData struct {
	Child *templ.Component
}

type EmptyData struct{}

type SquareHeadData EmptyData
type SquareBodyData EmptyData

func SquareBody(struct{}) *SquareBodyData {
	return &SquareBodyData{}
}

func SquareHead(struct{}) *SquareHeadData {
	return &SquareHeadData{}
}

type ProfileBodyCtr struct {
	R     *http.Request
	Child *templ.Component
}

func ProfileBody(c ProfileBodyCtr) *ProfileBodyData {
	username := c.R.PathValue("username")
	return &ProfileBodyData{
		username,
		c.Child,
	}
}

type ProfileHeadCtr struct {
	Child *templ.Component
}

func ProfileHead(c ProfileHeadCtr) *ProfileHeadData {
	return &ProfileHeadData{
		c.Child,
	}
}
