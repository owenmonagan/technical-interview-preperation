package graph

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func (graph *Graph) DFS(v *Vertex) {
	s := stack.New()
	graph.resetVisted()
	v.dfs(s)
	for s.Len() != 0 {
		fmt.Printf("%+v ", s.Pop())
	}
}

func (v *Vertex) dfs(s *stack.Stack) {
	if v.visted {
		return
	}
	s.Push(v.ID)
	v.visted = true
	current := v.edges
	for current != nil {
		current.to.dfs(s)
		current = current.next
	}
}

func (graph *Graph) BFS(v *Vertex) {
	queue := []string{}
	graph.resetVisted()
	fmt.Println(v.bfs(queue))
}

func (v *Vertex) bfs(queue []string) []string {
	if v.visted {
		return queue
	}
	queue = append(queue, v.ID)
	v.visted = true
	current := v.edges
	for current != nil {
		queue = current.to.bfs(queue)
		current = current.next
	}
	return queue
}
