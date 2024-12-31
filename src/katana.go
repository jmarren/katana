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
type TemplBodyFunc[S any] func(S) templ.Component
type TemplHeadFunc func() templ.Component
type ComponentFunc[T any] func(T) *Component

func NewComponent[T any, S any](ctr ComponentConstructor[T, S], bodyFunc TemplBodyFunc[S], headFunc TemplHeadFunc) ComponentFunc[T] {
	return func(t T) *Component {
		d := ctr(t)
		body := bodyFunc(d)
		head := headFunc()
		return &Component{
			&body,
			&head,
		}
	}
}

type ParentTemplate TemplBodyFunc[*Component]
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
