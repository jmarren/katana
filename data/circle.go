package data

type CircleBodyData struct {
	Color string
}
type CircleHeadData EmptyData

func CircleHead(d CircleHeadData) *CircleHeadData {
	return &CircleHeadData{}
}

func CircleBody(d CircleBodyData) *CircleBodyData {
	return &CircleBodyData{
		Color: d.Color,
	}
}
