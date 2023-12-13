package main

import (
	"reflect"
	"testing"
)

func Test_secondGreaterElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{nums: []int{2, 4, 0, 9, 6}},
			want: []int{9, 6, 6, -1, -1},
		},
		{
			name: "Example2",
			args: args{nums: []int{3, 3}},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondGreaterElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secondGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
