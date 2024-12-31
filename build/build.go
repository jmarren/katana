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
	setData()
	setHead(INode)
	setBody(INode)
	getName() string
	getHead() *templ.Component
	getBody() *templ.Component
	Render(w http.ResponseWriter, r *http.Request)
}

type SetHeadFunc func(string)
type TemplateFunc func(*templ.Component) templ.Component

// func CreateSetHead(node INode, headFunc TemplateFunc) SetHeadFunc  {
// 	return func(name string) INode  {
// 		child, err := GetNodeFactory(name)
// 		if err != nil {
// 			fmt.Printf("error:%s", err)
// 		}
// 		node.setHead(child.getName())
// 		head := headFunc(child.getHead())
// 	}
// }
//
//

// func (i *Index) setHead(name string) {
// 	child, err := GetNodeFactory(name)
// 	if err != nil {
// 		fmt.Printf("error:%s", err)
// 	}
// 	child.setHead(child.getName())
// 	head := templates.IndexHead(child.getHead())
// 	i.Head = &head
// }

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
