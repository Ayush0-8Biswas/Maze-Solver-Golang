package bin

type node struct {
	state
	action
	Parent *node
}

func Node(State state, Parent *node, Action action) *node {
	self := new(node)
	self.state = State
	self.Parent = Parent
	self.action = Action
	return self
}

type state struct {
	x int
	y int
}

type action string
