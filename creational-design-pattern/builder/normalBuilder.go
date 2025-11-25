package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() IBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden door"
}

func (n *NormalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}
