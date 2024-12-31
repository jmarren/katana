package build

import (
	// "fmt"
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	"net/http"
)

type About struct {
	Head *templ.Component
	Body *templ.Component
}

func (i *About) getName() string {
	return "about"
}

func (i *About) setHead(name string) {
	// child, err := GetNodeFactory(name)
	// if err != nil {
	// 	fmt.Printf("error:%s", err)
	// } else {
	// 	child.setHead(child.getName())
	// }
	head := templates.AboutHead()
	i.Head = &head
}

func (i *About) setBody(name string) {
	// child, err := GetNodeFactory(name)
	// if err != nil {
	// 	fmt.Printf("error:%s", err)
	// } else {
	// 	child.setBody(child.getName())
	// }
	body := templates.AboutBody()
	i.Body = &body
}

func (i *About) getHead() *templ.Component {
	return i.Head
}

func (i *About) getBody() *templ.Component {
	return i.Body
}

func (i *About) Render(w http.ResponseWriter, r *http.Request) {
	i.setBody("")
	i.setHead("")
	body := i.getBody()
	head := i.getHead()
	component := templates.Base(head, body)
	component.Render(r.Context(), w)
}
