package main

import "testing"

func Test_maxRemovals(t *testing.T) {
	type args struct {
		source        string
		pattern       string
		targetIndices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				source:        "abbaa",
				pattern:       "aba",
				targetIndices: []int{0, 1, 2},
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				source:        "bcda",
				pattern:       "d",
				targetIndices: []int{0, 3},
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				source:        "dda",
				pattern:       "dda",
				targetIndices: []int{0, 1, 2},
			},
			want: 0,
		},
		{
			name: "Example4",
			args: args{
				source:        "yeyeykyded",
				pattern:       "yeyyd",
				targetIndices: []int{0, 2, 3, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRemovals(tt.args.source, tt.args.pattern, tt.args.targetIndices); got != tt.want {
				t.Errorf("maxRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
