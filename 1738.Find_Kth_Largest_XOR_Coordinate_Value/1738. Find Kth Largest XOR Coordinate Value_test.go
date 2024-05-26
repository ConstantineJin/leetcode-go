package main

import "testing"

func Test_kthLargestValue(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				matrix: [][]int{{5, 2}, {1, 6}},
				k:      1,
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				matrix: [][]int{{5, 2}, {1, 6}},
				k:      2,
			},
			want: 5,
		},
		{
			name: "Example3",
			args: args{
				matrix: [][]int{{5, 2}, {1, 6}},
				k:      3,
			},
			want: 4,
		},
		{
			name: "Example4",
			args: args{
				matrix: [][]int{{5, 2}, {1, 6}},
				k:      4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
