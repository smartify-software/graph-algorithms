package graph_algorithms

type Paths struct {
	Dfs Dfs
}

func (p *Paths) HasPathTo(source, destination string) bool {
	visited := make(map[string]bool)
	p.Dfs.Dfs(source)
	return visited[destination]
}
