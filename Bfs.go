package graph_algorithms

type Bfs struct {
	Graph      Graph
	Queue      []string
	Visited    map[string]bool
	PathsGraph Graph
}

func (b Bfs) Bfs(
	source string,
) {
	b.Queue = append(b.Queue, source)
	b.Visited[source] = true
	for len(b.Queue) > 0 {
		vertex := b.Queue[0]
		b.Queue = b.Queue[1:]
		for _, w := range b.Graph[vertex] {
			if !b.Visited[w] {
				b.Visited[w] = true
				b.PathsGraph.AddDirectedEdge(vertex, w)
				b.Queue = append(b.Queue, w)
			}
		}
	}
}

func (b Bfs) Reset() {
	b.Queue = []string{}
	b.Visited = make(map[string]bool)
	b.PathsGraph = Graph{}
}
