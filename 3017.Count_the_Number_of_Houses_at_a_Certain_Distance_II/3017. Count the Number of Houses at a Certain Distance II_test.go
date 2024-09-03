package main

import (
	"reflect"
	"testing"
)

func Test_countOfPairs(t *testing.T) {
	type args struct {
		n int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Example1",
			args: args{
				n: 3,
				x: 1,
				y: 3,
			},
			want: []int64{6, 0, 0},
		},
		{
			name: "Example2",
			args: args{
				n: 5,
				x: 2,
				y: 4,
			},
			want: []int64{10, 8, 2, 0, 0},
		},
		{
			name: "Example3",
			args: args{
				n: 4,
				x: 1,
				y: 1,
			},
			want: []int64{6, 4, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfPairs(tt.args.n, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
