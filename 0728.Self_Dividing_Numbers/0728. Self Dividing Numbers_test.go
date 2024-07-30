package main

import (
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				left:  1,
				right: 22,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			name: "Example2",
			args: args{
				left:  47,
				right: 85,
			},
			want: []int{48, 55, 66, 77},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selfDividingNumbers(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selfDividingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
