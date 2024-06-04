package main

import (
	"reflect"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	type args struct {
		rains []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{rains: []int{1, 2, 3, 4}},
			want: []int{-1, -1, -1, -1},
		},
		{
			name: "Example2",
			args: args{rains: []int{1, 2, 0, 0, 2, 1}},
			want: []int{-1, -1, 2, 1, -1, -1},
		},
		{
			name: "Example3",
			args: args{rains: []int{1, 2, 0, 1, 2}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avoidFlood(tt.args.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
