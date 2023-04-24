package graph_algorithms

type TopologicalSort struct {
	Graph   Graph
	Stack   []string
	Visited map[string]bool
}

func (t TopologicalSort) Sort() {
	for key := range t.Graph {
		if !t.Visited[key] {
			t.Dfs(key)
		}
	}
}

func (t TopologicalSort) Dfs(
	key string,
) {
	t.Visited[key] = true
	for _, w := range t.Graph[key] {
		if !t.Visited[w] {
			t.Dfs(w)
		}
	}
	t.Stack = append([]string{key}, t.Stack...)
}

func (t TopologicalSort) Reset() {
	t.Stack = []string{}
}
