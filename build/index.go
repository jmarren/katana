package build

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	"net/http"
)

type Index struct {
	Head *templ.Component
	Body *templ.Component
}

func (i *Index) setHead(name string) {
	child, err := GetNodeFactory(name)
	if err != nil {
		fmt.Printf("error:%s", err)
	}
	child.setHead(child.getName())
	head := templates.IndexHead(child.getHead())
	i.Head = &head
}

func (i *Index) setBody(name string) {
	child, err := GetNodeFactory(name)
	if err != nil {
		fmt.Printf("error:%s", err)
	}
	child.setBody(child.getName())
	body := templates.IndexBody(child.getBody())
	i.Body = &body
}

func (i *Index) getHead() *templ.Component {
	return i.Head
}

func (i *Index) getBody() *templ.Component {
	return i.Body
}

func (i *Index) getName() string {
	return "index"
}

func (i *Index) Render(w http.ResponseWriter, r *http.Request) {
	page := r.PathValue("page")
	child, err := GetNodeFactory(page)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	i.setBody(child.getName())
	i.setHead(child.getName())
	body := i.getBody()
	head := i.getHead()
	component := templates.Base(head, body)
	component.Render(r.Context(), w)
}
