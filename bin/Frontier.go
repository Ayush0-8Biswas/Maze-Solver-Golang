package bin

type stackFrontier struct {
	Frontier []*node
}

func StackFrontier() *stackFrontier {
	self := new(stackFrontier)
	self.Frontier = make([]*node, 0)
	return self
}

func (self *stackFrontier) Add(Node *node) {
	self.Frontier = append(self.Frontier, Node)
}

func (self *stackFrontier) ContainsState(State state) bool {
	for _, nd := range self.Frontier {
		if nd.state.x == State.x && nd.state.y == State.y {
			return true
		}
	}
	return false
}

func (self *stackFrontier) Empty() bool {
	return len(self.Frontier) == 0
}

func (self *stackFrontier) Remove() *node {
	if self.Empty() {
		panic("empty self")
	} else {
		nd := self.Frontier[len(self.Frontier)-1]
		self.Frontier = self.Frontier[:len(self.Frontier)-1]
		return nd
	}
}

type queueFrontier struct {
	stackFrontier
}

func QueueFrontier() *queueFrontier {
	self := new(queueFrontier)
	self.Frontier = make([]*node, 0)
	return self
}

func (self *queueFrontier) Remove() *node {
	if self.Empty() {
		panic("empty self")
	} else {
		nd := self.Frontier[0]
		self.Frontier = self.Frontier[1:]
		return nd
	}
}
