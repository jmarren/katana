package build

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	"net/http"
)

type Index struct {
	Title string
	Component
}

func (i *Index) setHead(child INode) {
	head := templates.IndexHead(child.getHead())
	i.Head = &head
}

func (i *Index) setBody(child INode) {
	body := templates.IndexBody(child.getBody())
	i.Body = &body
}

func (i *Index) setData() {
	i.Title = "index page"
}

func (i *Index) setChild(child INode) {
	child.setData()
	i.setData()
	i.setHead(child)
	i.setBody(child)
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
	child.setData()
	child.setBody(nil)
	child.setHead(nil)
	i.setBody(child)
	i.setHead(child)
	body := i.getBody()
	head := i.getHead()
	component := templates.Base(head, body)
	component.Render(r.Context(), w)
}
