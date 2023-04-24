package main

import api "bitbucket.org/smartify.software/graph-algorithms"

func main() {
	graph := api.Graph{
		"A": []string{"B", "C"},
		"B": []string{"A", "C", "D"},
		"C": []string{"A", "B", "D", "E"},
		"F": []string{},
	}

	dfs := api.Dfs{
		Graph:      graph,
		Visited:    make(map[string]bool),
		PathsGraph: api.Graph{},
	}

	bfs := api.Bfs{
		Graph:      graph,
		Queue:      []string{},
		Visited:    make(map[string]bool),
		PathsGraph: api.Graph{},
	}
	dfs.Dfs("B")
	bfs.Bfs("B")

	file := api.PrettyPrintGraphvizFile{PrettyPintGraphviz: api.PrettyPintGraphviz{Graph: graph}}
	err := file.Write("graph.dot")
	if err != nil {
		panic(err)
	}

	file = api.PrettyPrintGraphvizFile{PrettyPintGraphviz: api.PrettyPintGraphviz{Graph: dfs.PathsGraph}}
	err = file.Write("dfs_paths_graph.dot")
	if err != nil {
		panic(err)
	}

	file = api.PrettyPrintGraphvizFile{PrettyPintGraphviz: api.PrettyPintGraphviz{Graph: bfs.PathsGraph}}
	err = file.Write("bfs_paths_graph.dot")
	if err != nil {
		panic(err)
	}

	connectedComponents := api.ConnectedComponents{Dfs: dfs}
	println(connectedComponents.Count())
}
