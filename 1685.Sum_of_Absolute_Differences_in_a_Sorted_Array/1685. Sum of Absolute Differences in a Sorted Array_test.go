package main

import (
	"reflect"
	"testing"
)

func Test_getSumAbsoluteDifferences(t *testing.T) {
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
			args: args{nums: []int{2, 3, 5}},
			want: []int{4, 3, 5},
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 4, 6, 8, 10}},
			want: []int{24, 15, 13, 15, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumAbsoluteDifferences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumAbsoluteDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
