package components

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildComponent() Component {
	// d.builder.Parent.setChildBodyFunc()
	// d.builder.Parent.setChildHeadFunc()
	d.builder.setParent()
	d.builder.setBody()
	d.builder.setHead()
	return d.builder.getComponent()
}
