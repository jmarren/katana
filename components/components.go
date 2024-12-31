package components

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	"net/http"
)

type IComponentData interface {
}

type Component struct {
	Head *templ.Component
	Body *templ.Component
}

type IComponent interface {
	Render()
}

func (c Component) Render(w http.ResponseWriter, r *http.Request) {
	full := templates.Base(c.Head, c.Body)
	full.Render(r.Context(), w)
}

func PrintHi() {
	fmt.Println("hi!")
}
