package main

import "testing"

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{
				quality: []int{10, 20, 5},
				wage:    []int{70, 50, 30},
				k:       2,
			},
			want: 105,
		},
		{
			name: "Example2",
			args: args{
				quality: []int{3, 1, 10, 10, 1},
				wage:    []int{4, 8, 2, 2, 7},
				k:       3,
			},
			want: 92.0 / 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostToHireWorkers(tt.args.quality, tt.args.wage, tt.args.k); got != tt.want {
				t.Errorf("mincostToHireWorkers() = %v, want %v", got, tt.want)
			}
		})
	}
}
