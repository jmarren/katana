package components

import (
	"github.com/a-h/templ"
	// "github.com/jmarren/katana/templates"
)

type ChildFunc func(*templ.Component) templ.Component

type BaseBuilder struct {
	Head          *templ.Component
	Body          *templ.Component
	ChildBodyFunc ChildFunc
	ChildHeadFunc ChildFunc
}

func newBaseBuilder() *BaseBuilder {
	return &BaseBuilder{}
}

func (b *BaseBuilder) getComponent() *Component {
	return &Component{
		b.Head,
		b.Body,
	}
}

// func (b *BaseBuilder) setBody() {
//
// 	body := templates.Base(nil)
// 	b.Body = &body
// }
//
// func (b *BaseBuilder) setHead() {
// 	head := templates.Head(nil)
// 	b.Head = &head
// }

// func (b *BaseBuilder) setChildBodyFunc() {
// 	childBodyFunc := templates.Base
// 	b.ChildBodyFunc = childBodyFunc
// }

// func (b *BaseBuilder) setChildHeadFunc() {
// 	childHeadFunc := templates.Head
// 	b.ChildHeadFunc = childHeadFunc
// }
