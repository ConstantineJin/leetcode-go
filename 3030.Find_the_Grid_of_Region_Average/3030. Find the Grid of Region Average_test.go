package main

import (
	"reflect"
	"testing"
)

func Test_resultGrid(t *testing.T) {
	type args struct {
		image     [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				image:     [][]int{{5, 6, 7, 10}, {8, 9, 10, 10}, {11, 12, 13, 10}},
				threshold: 3,
			},
			want: [][]int{{9, 9, 9, 9}, {9, 9, 9, 9}, {9, 9, 9, 9}},
		},
		{
			name: "Example2",
			args: args{
				image:     [][]int{{10, 20, 30}, {15, 25, 35}, {20, 30, 40}, {25, 35, 45}},
				threshold: 12,
			},
			want: [][]int{{25, 25, 25}, {27, 27, 27}, {27, 27, 27}, {30, 30, 30}},
		},
		{
			name: "Example3",
			args: args{
				image:     [][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}},
				threshold: 1,
			},
			want: [][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultGrid(tt.args.image, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
