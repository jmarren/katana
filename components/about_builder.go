package components

import (
	// "github.com/a-h/templ"
	"github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
	// "net/http"
)

type AboutBuilder struct {
	Parent IParentBuilder
	Head   *templ.Component
	Body   *templ.Component
}

func newAboutBuilder() IBuilder {
	return &AboutBuilder{}
}

// func (b *AboutBuilder) setData() {
// 	b.setData()
// }

func (b *AboutBuilder) getComponent() Component {
	head := b.Parent.getChildHeadFunc()(b.Head)
	body := b.Parent.getChildBodyFunc()(b.Body)
	return Component{
		&head,
		&body,
	}
}

func (b *AboutBuilder) setParent() {
	b.Parent = &IndexBuilder{}
	b.Parent.setChildBodyFunc()
	b.Parent.setChildHeadFunc()
}

func (b *AboutBuilder) setHead() {
	head := templates.AboutHead()
	b.Head = &head
}

func (b *AboutBuilder) setBody() {
	body := templates.AboutBody()
	b.Body = &body
}
