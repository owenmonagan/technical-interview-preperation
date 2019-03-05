package graph

import (
	"testing"
)

func TestNew(t *testing.T) {
	graph := New()
	if graph.vertex != nil {
		t.Error("should of been empty")
	}
}

func TestAddVertex(t *testing.T) {
	graph := New()
	idA := "a"
	graph.addVertex(&Vertex{ID: idA})
	if graph.vertex.ID != idA {
		t.Error("didn't add the vertext")
	}

	idB := "b"
	graph.addVertex(&Vertex{ID: idB})

	if graph.vertex.next.ID != idB {
		t.Error("didn't add the vertext")
	}

	idC := "c"
	graph.addVertex(&Vertex{ID: idC})
	if graph.vertex.next.next.ID != idC {
		t.Error("didn't add the vertext")
	}
	if graph.addVertex(&Vertex{ID: idC}) {
		t.Error("should not of been able to add an existing vertex")
	}
}

func TestSearchVertex(t *testing.T) {
	graph := New()
	idA := "a"
	if graph.get(idA) != nil {
		t.Error("should of gotten nil")
	}
	graph.addVertex(&Vertex{ID: idA})

	idB := "b"
	graph.addVertex(&Vertex{ID: idB})

	idC := "c"
	graph.addVertex(&Vertex{ID: idC})

	if graph.get(idC).ID != idC {
		t.Error("should of gotten id c")
	}
	if graph.get(idA).ID != idA {
		t.Error("should of gotten id c")
	}
}

func TestAddEdge(t *testing.T) {
	graph := New()

	idA := "a"
	idB := "b"
	idC := "c"
	idD := "D"

	if graph.addEdge(idA, &Edge{to: graph.get(idB)}) {
		t.Error("should not of been able to add an edge to a non existant vertex")
	}

	graph.addVertex(&Vertex{ID: idA})
	graph.addVertex(&Vertex{ID: idB})
	graph.addVertex(&Vertex{ID: idC})
	graph.addVertex(&Vertex{ID: idD})

	graph.addEdge(idA, &Edge{to: graph.get(idC)})
	if graph.get(idA).edges.to.ID != idC {
		t.Error("should of gotten this vertex")
	}
	graph.addEdge(idA, &Edge{to: graph.get(idB)})
	if graph.get(idA).edges.next.to.ID != idB {
		t.Error("should of gotten this vertex")
	}
	graph.addEdge(idA, &Edge{to: graph.get(idD)})
	if graph.get(idA).edges.next.next.to.ID != idD {
		t.Error("should of gotten this vertex")
	}

	graph.addEdge(idC, &Edge{to: graph.get(idB)})
	if graph.addEdge(idC, &Edge{to: graph.get(idB)}) {
		t.Error("should not of been able to add an existing edge")
	}
}

func TestReachable(t *testing.T) {
	graph := New()
	idA := "a"
	idB := "b"
	idC := "c"
	idD := "D"
	idZ := "Z"
	idX := "X"

	graph.addVertex(&Vertex{ID: idA})
	graph.addVertex(&Vertex{ID: idB})
	graph.addVertex(&Vertex{ID: idC})
	graph.addVertex(&Vertex{ID: idD})
	graph.addVertex(&Vertex{ID: idZ})
	graph.addVertex(&Vertex{ID: idX})

	if graph.isReachable(graph.get(idA), graph.get(idB)) {
		t.Error("should not be reachable")
	}
	graph.addEdge(idA, &Edge{to: graph.get(idZ)})

	graph.addEdge(idA, &Edge{to: graph.get(idX)})
	if !graph.isReachable(graph.get(idA), graph.get(idX)) {
		t.Error("should be reachable")
	}
	graph.addEdge(idA, &Edge{to: graph.get(idC)})
	graph.addEdge(idC, &Edge{to: graph.get(idB)})
	graph.addEdge(idB, &Edge{to: graph.get(idD)})
	if !graph.isReachable(graph.get(idA), graph.get(idD)) {
		t.Error("should be reachable")
	}
}
