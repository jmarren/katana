package data

import (
	// "github.com/jmarren/katana/src"
	"net/http"
	// "github.com/jmarren/katana/templates"
)

type ProfileData struct {
	Username string
	Square   SquareData
}

type EmptyData struct{}

type SquareData EmptyData

func Square(EmptyData) *SquareData {
	return &SquareData{}
}

func Profile(r *http.Request) *ProfileData {
	username := r.PathValue("username")
	return &ProfileData{
		username,
		SquareData{},
	}
}
