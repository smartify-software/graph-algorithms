package graph_algorithms

type PrettyPintGraphviz struct {
	Graph Graph
}

func (g PrettyPintGraphviz) PrettyPrint() string {
	dotString := "digraph G {\n"
	for key, value := range g.Graph {
		if len(value) == 0 {
			dotString += "\t" + key + "\n"
		}
		for _, v := range value {
			dotString += "\t" + key + " -> " + v + "\n"
		}
	}
	dotString += "}"
	return dotString
}
