package build

import (
	// "fmt"
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	"net/http"
)

type About struct {
	Component
	Number int
}

func (i *About) getName() string {
	return "about"
}

func (i *About) setData() {
	i.Number = 10
}

func (i *About) setHead(child INode) {
	head := templates.AboutHead()
	i.Head = &head
}

func (i *About) setBody(child INode) {
	body := templates.AboutBody(i.Number)
	i.Body = &body
}

func (i *About) getHead() *templ.Component {
	return i.Head
}

func (i *About) getBody() *templ.Component {
	return i.Body
}

func (i *About) Render(w http.ResponseWriter, r *http.Request) {
	i.setData()
	i.setBody(nil)
	i.setHead(nil)
	body := i.getBody()
	head := i.getHead()
	component := templates.Base(head, body)
	component.Render(r.Context(), w)
}
