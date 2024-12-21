package main

import (
	"reflect"
	"testing"
)

func Test_sortTheStudents(t *testing.T) {
	type args struct {
		score [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				score: [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}},
				k:     2,
			},
			want: [][]int{{7, 5, 11, 2}, {10, 6, 9, 1}, {4, 8, 3, 15}},
		},
		{
			name: "Example2",
			args: args{
				score: [][]int{{3, 4}, {5, 6}},
				k:     0,
			},
			want: [][]int{{5, 6}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTheStudents(tt.args.score, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTheStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
