package graph_algorithms

type Graph map[string][]string

func (g Graph) AddVertex(vertex string) {
	g[vertex] = []string{}
}

func (g Graph) AddDirectedEdge(vertex, edge string) {
	g[vertex] = append(g[vertex], edge)
}

func (g Graph) AddEdge(vertex, edge string) {
	g[vertex] = append(g[vertex], edge)
	g[edge] = append(g[edge], vertex)
}
