package main

import api "bitbucket.org/smartify.software/graph-algorithms"

func main() {
	graph := api.Graph{
		"A": []string{"B", "C"},
		"B": []string{"A", "C", "D"},
		"C": []string{"A", "B", "D", "E"},
	}

	paths := api.Paths{Graph: graph}
	pathsGraph := api.Graph{}
	paths.Dfs("B", make(map[string]bool), pathsGraph)

	file := api.PrettyPrintGraphvizFile{PrettyPintGraphviz: api.PrettyPintGraphviz{Graph: graph}}
	err := file.Write("graph.dot")
	if err != nil {
		panic(err)
	}

	file = api.PrettyPrintGraphvizFile{PrettyPintGraphviz: api.PrettyPintGraphviz{Graph: pathsGraph}}
	err = file.Write("paths_graph.dot")
	if err != nil {
		panic(err)
	}
}
