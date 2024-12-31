package components

// import "net/http"

type IParentBuilder interface {
	setChildBodyFunc()
	setChildHeadFunc()
	getChildHeadFunc() ChildFunc
	getChildBodyFunc() ChildFunc
}

type IBuilder interface {
	// IParentBuilder
	setParent()
	setHead()
	setBody()
	getComponent() Component
}

func getBuilder(builderType string) IBuilder {
	switch builderType {
	case "about":
		return newAboutBuilder()
	}
	return nil
}
