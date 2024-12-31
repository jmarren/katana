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
func (i *Contact) setHead(name string) {
	head := templates.ContactHead()
	i.Head = &head
}

func (i *Contact) setBody(name string) {
	body := templates.ContactBody()
	i.Body = &body
}

func (i *Contact) getHead() *templ.Component {
	return i.Head
}

func (i *Contact) getBody() *templ.Component {
	return i.Body
}

func (i *Contact) Render(w http.ResponseWriter, r *http.Request) {
	i.setBody("")
	i.setHead("")
	body := i.getBody()
	head := i.getHead()
	component := templates.Base(head, body)
	component.Render(r.Context(), w)
}
