package build

import (
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	"net/http"
)

type Contact Component

func (i *Contact) getName() string {
	return "contact"
}
func (i *Contact) setHead(INode) {
	head := templates.ContactHead()
	i.Head = &head
}

func (i *Contact) setBody(INode) {
	body := templates.ContactBody()
	i.Body = &body
}

func (i *Contact) setData() {
}

func (i *Contact) getHead() *templ.Component {
	return i.Head
}

func (i *Contact) getBody() *templ.Component {
	return i.Body
}

func (i *Contact) Render(w http.ResponseWriter, r *http.Request) {
	i.setData()
	i.setBody(nil)
	i.setHead(nil)
	body := i.getBody()
	head := i.getHead()
	component := templates.Base(head, body)
	component.Render(r.Context(), w)
}
