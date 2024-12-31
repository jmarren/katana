package src

import (
	"github.com/a-h/templ"
	"net/http"
)

type Component struct {
	Body *templ.Component
	Head *templ.Component
}

type ComponentConstructor[T any, S any] func(T) S
type TemplFunc[S any] func(S) templ.Component

// type TemplFunc func() templ.Component
type ComponentFunc[T any, U any] func(T, U) *Component

// type TemplFunc[T any] func(T) templ.Component

// Takes a data constructor, a bodyFunc and a headFunc and returns a function that will accept some data,
// construct it using the constructor, inject the result into the body func, and return a Component
func NewComponent[T any, S any, U any, V any](ctr ComponentConstructor[T, S], headCtr ComponentConstructor[U, V], bodyFunc TemplFunc[S], headFunc TemplFunc[V]) ComponentFunc[T, U] {
	return func(t T, u U) *Component {
		bodyData := ctr(t)
		headData := headCtr(u)
		body := bodyFunc(bodyData)
		head := headFunc(headData)
		return &Component{
			&body,
			&head,
		}
	}
}

// func NewParent[T any, S any](ctr ComponentConstructor[T, S], bodyFunc TemplBodyFunc[S], headFunc TemplHeadFunc) ComponentFunc[T] {
// 	return func(t T) *Component {
// 		d := ctr(t)
// 		body := bodyFunc(d)
// 		head := headFunc()
// 		return &Component{
// 			&body,
// 			&head,
// 		}
// 	}
// }

// I need a function that will take a constructor that takes some data, a constructor, and a Component that
// func NewParent(child Component, bodyFunc

// type ParentConstructor[T any] func(T, *templ.Component, *templ.Component)
// func

// type WrapperComponent ComponentFunc[[]*templ.Component]
type ParentTemplate TemplFunc[*Component]
type WrapperTemplate func(*templ.Component) templ.Component

func (c *Component) Render(w http.ResponseWriter, r *http.Request, baseFunc ParentTemplate) {
	full := baseFunc(c)
	full.Render(r.Context(), w)
}

func (c *Component) Wrap(headWrapper WrapperTemplate, bodyWrapper WrapperTemplate) *Component {
	head := headWrapper(c.Head)
	body := bodyWrapper(c.Body)
	return &Component{
		&body,
		&head,
	}
}
