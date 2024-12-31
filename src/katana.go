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

type ComponentFuncCtr[T any, S any, U any, V any] struct {
	BodyCtr  ComponentConstructor[T, S]
	HeadCtr  ComponentConstructor[U, V]
	BodyFunc TemplFunc[S]
	HeadFunc TemplFunc[V]
}

type EltCtr[T any, S any] struct {
	Ctr       ComponentConstructor[T, S]
	TemplFunc TemplFunc[S]
}

// type TemplFunc[T any] func(T) templ.Component

// Takes a data constructor, a bodyFunc and a headFunc and returns a function that will accept some data,
// construct it using the constructor, inject the result into the body func, and return a Component
func NewComponent[T any, S any, U any, V any](bodyElt EltCtr[T, S], headElt EltCtr[U, V]) ComponentFunc[T, U] {
	return func(t T, u U) *Component {
		bodyData := bodyElt.Ctr(t)
		headData := headElt.Ctr(u)
		body := bodyElt.TemplFunc(bodyData)
		head := headElt.TemplFunc(headData)
		return &Component{
			&body,
			&head,
		}
	}
}

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
