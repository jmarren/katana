package build2

import (
	"github.com/a-h/templ"
	// "github.com/jmarren/katana/templates"
	"net/http"
)

type Component struct {
	Head *templ.Component
	Body *templ.Component
}

type IComponent interface {
	GetHead() *templ.Component
	GetBody() *templ.Component
}

type BodyFunc[T any] func(T) templ.Component

// type HeadFunc func() templ.Component
type DataFunc[T any] func(*http.Request) T
type Head templ.Component

type ComponentCtr[T any] struct {
	BodyFunc[T]
	Head
	DataFunc[T]
}

type IndexData struct {
	Title  string
	Number int
}

// var index = &ComponentCtr[*IndexData]{
// 	DataFunc: func(r *http.Request) *IndexData {
// 		return &IndexData{
// 			Title:  "Index Page",
// 			Number: 10,
// 		}
// 	},
// 	BodyFunc: func(iData *IndexData) {
//
// 	},
// 	Head: templates.AboutHead(),
// }

// type PageConstructor[T any] func(*http.Request) T
// type ComponentCtr[T any]

// type [T any]ComponentCtr struct {
//
// }

// type Page struct {
// 	Head *templ.Component
// }
