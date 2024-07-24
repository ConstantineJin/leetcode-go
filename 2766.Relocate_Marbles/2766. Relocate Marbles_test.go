package main

import (
	"reflect"
	"testing"
)

func Test_relocateMarbles(t *testing.T) {
	type args struct {
		nums     []int
		moveFrom []int
		moveTo   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:     []int{1, 6, 7, 8},
				moveFrom: []int{1, 7, 2},
				moveTo:   []int{2, 9, 5},
			},
			want: []int{5, 6, 8, 9},
		},
		{
			name: "Example2",
			args: args{
				nums:     []int{1, 1, 3, 3},
				moveFrom: []int{1, 3},
				moveTo:   []int{2, 2},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relocateMarbles(tt.args.nums, tt.args.moveFrom, tt.args.moveTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relocateMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
