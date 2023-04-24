package graph_algorithms

type Dfs struct {
	Graph      Graph
	Visited    map[string]bool
	PathsGraph Graph
}

func (d Dfs) Dfs(
	source string,
) {
	d.Visited[source] = true
	for _, w := range d.Graph[source] {
		if !d.Visited[w] {
			d.PathsGraph.AddDirectedEdge(source, w)
			d.Dfs(w)
		}
	}
}

func (d Dfs) Reset() {
	d.Visited = make(map[string]bool)
	d.PathsGraph = Graph{}
}
