package graph_algorithms

type ConnectedComponents struct {
	Dfs Dfs
}

func (c *ConnectedComponents) Count() int {
	visited := make(map[string]bool)
	count := 0
	for key := range c.Dfs.Graph {
		if !visited[key] {
			c.Dfs.Dfs(key)
			count++
		}
	}
	return count
}
