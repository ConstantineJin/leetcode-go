package main

import "testing"

func Test_findChampion(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name      string
		args      args
		wantChamp int
	}{
		{
			name:      "Example1",
			args:      args{n: 3, edges: [][]int{{0, 1}, {1, 2}}},
			wantChamp: 0,
		},
		{
			name:      "Example2",
			args:      args{n: 4, edges: [][]int{{0, 2}, {1, 3}, {1, 2}}},
			wantChamp: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChamp := findChampion(tt.args.n, tt.args.edges); gotChamp != tt.wantChamp {
				t.Errorf("findChampion() = %v, want %v", gotChamp, tt.wantChamp)
			}
		})
	}
}
