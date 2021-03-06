package graph

type Graph struct {
	vertex *Vertex
}

type Vertex struct {
	ID     string
	next   *Vertex
	edges  *Edge
	visted bool
}

type Edge struct {
	len  int
	to   *Vertex
	next *Edge
}

func New() (graph Graph) {
	graph = Graph{}
	return
}

func (graph *Graph) addVertex(v *Vertex) bool {
	if *graph == (Graph{}) {
		graph.vertex = v
		return true
	}
	current := graph.vertex
	for {
		if current.ID == v.ID {
			return false
		} else if current.next == nil {
			current.next = v
			return true
		} else {
			current = current.next
		}
	}
}

func (graph *Graph) get(id string) *Vertex {
	current := graph.vertex
	for {
		if current == nil {
			return current
		} else if current.ID == id {
			return current
		} else {
			current = current.next
		}
	}
}

func (graph *Graph) addEdge(id string, e *Edge) bool {
	vertex := graph.get(id)
	if vertex == nil {
		return false
	}
	if vertex.edges == nil {
		vertex.edges = e
		return true
	}
	current := vertex.edges
	for {
		if current.to.ID == e.to.ID {
			return false
		} else if current.next == nil {
			current.next = e
			return true
		} else {
			current = current.next
		}
	}
}

func (graph *Graph) resetVisted() {
	current := graph.vertex
	for current != nil {
		current.visted = false
		current = current.next
	}
}

func (graph *Graph) isReachable(s, e *Vertex) bool {
	graph.resetVisted()
	return s.reachable(e.ID)
}

func (v *Vertex) reachable(endID string) bool {
	if v.visted {
		return false
	}
	if v.ID == endID {
		return true
	}
	v.visted = true
	current := v.edges
	for {
		if current == nil {
			return false
		}
		if current.to.reachable(endID) {
			return true
		}
		current = current.next
	}
}
