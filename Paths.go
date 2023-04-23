package graph_algorithms

type Paths struct {
	Graph Graph
}

func (p *Paths) Dfs(
	source string,
	visited map[string]bool,
	pathsGraph Graph,
) {
	visited[source] = true
	for _, w := range p.Graph[source] {
		if !visited[w] {
			pathsGraph.AddDirectedEdge(source, w)
			p.Dfs(w, visited, pathsGraph)
		}
	}
}
