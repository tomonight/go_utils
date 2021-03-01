package graph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		data []string
	}
	data := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		{name: "success", args: args{data: data}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGraph(tt.args.data)
			got.InsertEdgeByPoint("1", "2", 1)
			got.InsertEdgeByPoint("2", "4", 1)
			got.InsertEdgeByPoint("2", "5", 1)
			got.InsertEdgeByPoint("4", "8", 1)
			got.InsertEdgeByPoint("5", "8", 1)
			got.InsertEdgeByPoint("1", "3", 1)
			got.InsertEdgeByPoint("3", "6", 1)
			got.InsertEdgeByPoint("3", "7", 1)
			got.InsertEdgeByPoint("6", "7", 1)
			//got.AddPoint("F")
			//got.InsertEdgeByPoint("A","F",1)
			got.ShowGraph()
			got.DFS()
			fmt.Println(got)
		})
	}
}
