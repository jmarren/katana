package components

import (
	// "github.com/a-h/templ"
	// "github.com/a-h/templ"
	"github.com/jmarren/katana/templates"
)

type IndexBuilder struct {
	// BaseBuilder
	// Head          *templ.Component
	// Body          *templ.Component
	ChildHeadFunc ChildFunc
	ChildBodyFunc ChildFunc
	Title         string
}

func newIndexBuilder() IParentBuilder {
	return &IndexBuilder{}
}

// func (b *IndexBuilder) getComponent() *Component {
// 	return &Component{
// 		b.Body,
// 		b.Head,
// 	}
// }

// func (b *IndexBuilder) getComponent() *Component {
// 	head := b.ChildHeadFunc(b.Head)
// 	body := b.ChildBodyFunc(b.Body)
// 	return &Component{
// 		&head,
// 		&body,
// 	}
// }

func (b *IndexBuilder) getChildBodyFunc() ChildFunc {
	return b.ChildBodyFunc
}

func (b *IndexBuilder) getChildHeadFunc() ChildFunc {
	return b.ChildHeadFunc
}

func (b *IndexBuilder) setChildBodyFunc() {
	childBodyFunc := templates.IndexBody
	b.ChildBodyFunc = childBodyFunc
}

func (b *IndexBuilder) setChildHeadFunc() {
	childHeadFunc := templates.IndexHead
	b.ChildBodyFunc = childHeadFunc
}

func (b *IndexBuilder) setTitle() {
	b.Title = "index"
}

func (b *IndexBuilder) setData() {
	b.setTitle()
}

// func (b *IndexBuilder) setHead() {
// 	head := templates.IndexHead()
// 	b.Head = &head
// }

// func (b *IndexBuilder) setBody() {
// 	head := templates.IndexBody(nil)
// 	b.Head = &head
// }
