package components

import (
	"fmt"
)

func Run() Component {
	indexBuilder := getBuilder("about")
	director := newDirector(indexBuilder)
	indexComponent := director.buildComponent()
	fmt.Printf("index component: %v\n", indexComponent)
	return indexComponent
}
