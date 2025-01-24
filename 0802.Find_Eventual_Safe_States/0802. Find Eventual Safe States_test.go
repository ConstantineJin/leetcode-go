package main

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name:    "Example1",
			args:    args{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
			wantAns: []int{2, 4, 5, 6},
		},
		{
			name:    "Example2",
			args:    args{graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}},
			wantAns: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("eventualSafeNodes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
