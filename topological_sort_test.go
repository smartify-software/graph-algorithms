package graph_algorithms

import "testing"

func TestTopologicalSort_Dfs(t1 *testing.T) {
	type fields struct {
		Graph Graph
		Stack []string
	}
	type args struct {
		key     string
		visited map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Graph: Graph{
					"0": {"1", "2"},
					"1": {"3"},
					"2": {"3"},
					"3": {"4"},
				},
				Stack: []string{},
			},
			args: args{
				key:     "0",
				visited: make(map[string]bool),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TopologicalSort{
				Graph:   tt.fields.Graph,
				Stack:   tt.fields.Stack,
				Visited: make(map[string]bool),
			}
			t.Dfs(tt.args.key)
			sort := TopologicalSort{
				Graph:   t.Graph,
				Stack:   t.Stack,
				Visited: t.Visited,
			}
			file := PrettyPrintGraphvizFile{PrettyPintGraphviz: PrettyPintGraphviz{sort.Graph}}
			file.PrettyPrint()
		})
	}
}
