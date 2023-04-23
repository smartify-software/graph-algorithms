package graph_algorithms

type PrettyPintGraphviz struct {
	Graph Graph
}

func (g PrettyPintGraphviz) PrettyPrint() string {
	dotString := "digraph G {\n"
	for key, value := range g.Graph {
		for _, v := range value {
			dotString += "\t" + key + " -> " + v + "\n"
		}
	}
	dotString += "}"
	return dotString
}
