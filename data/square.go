package data

import (
// "github.com/jmarren/katana/src"
// "net/http"

// "github.com/a-h/templ"
)

type EmptyData struct{}

type SquareBodyData struct {
	Color string
}
type SquareHeadData EmptyData

func SquareHead(d SquareHeadData) *SquareHeadData {
	return &SquareHeadData{}
}

func SquareBody(d SquareBodyData) *SquareBodyData {
	return &SquareBodyData{
		Color: d.Color,
	}
}
