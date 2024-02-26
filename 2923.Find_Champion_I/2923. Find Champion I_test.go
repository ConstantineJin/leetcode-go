package main

import "testing"

func Test_findChampion(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name         string
		args         args
		wantChampion int
	}{
		{
			name:         "Example1",
			args:         args{grid: [][]int{{0, 1}, {0, 0}}},
			wantChampion: 0,
		},
		{
			name:         "Example2",
			args:         args{grid: [][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}}},
			wantChampion: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChampion := findChampion(tt.args.grid); gotChampion != tt.wantChampion {
				t.Errorf("findChampion() = %v, want %v", gotChampion, tt.wantChampion)
			}
		})
	}
}
