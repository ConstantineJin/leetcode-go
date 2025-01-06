package main

import (
	"reflect"
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		boxes string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{boxes: "110"},
			want: []int{1, 1, 3},
		},
		{
			name: "Example2",
			args: args{boxes: "001011"},
			want: []int{11, 8, 5, 4, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
