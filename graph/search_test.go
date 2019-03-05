package graph

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	graph := New()
	idA := "A"
	idB := "B"
	idC := "C"
	idD := "D"
	idZ := "Z"
	idX := "X"

	graph.addVertex(&Vertex{ID: idA})
	graph.addVertex(&Vertex{ID: idB})
	graph.addVertex(&Vertex{ID: idC})
	graph.addVertex(&Vertex{ID: idD})
	graph.addVertex(&Vertex{ID: idZ})
	graph.addVertex(&Vertex{ID: idX})
	graph.addEdge(idA, &Edge{to: graph.get(idZ)})
	graph.addEdge(idA, &Edge{to: graph.get(idX)})
	graph.addEdge(idX, &Edge{to: graph.get(idA)})

	graph.addEdge(idA, &Edge{to: graph.get(idC)})
	graph.addEdge(idC, &Edge{to: graph.get(idB)})
	graph.addEdge(idB, &Edge{to: graph.get(idD)})
	graph.DFS(graph.get(idA))
	fmt.Println()
	graph.BFS(graph.get(idA))
}
