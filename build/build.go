package build

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
)

type Component struct {
	Head *templ.Component
	Body *templ.Component
}

type Node struct {
	INode
}

type INode interface {
	setHead(string)
	setBody(string)
	getName() string
	getHead() *templ.Component
	getBody() *templ.Component
	Render(w http.ResponseWriter, r *http.Request)
}

func GetNodeFactory(name string) (INode, error) {
	if name == "index" {
		return &Index{}, nil
	}
	if name == "about" {
		return &About{}, nil
	}

	if name == "contact" {
		return &Contact{}, nil
	}
	return nil, fmt.Errorf("not found")
}
