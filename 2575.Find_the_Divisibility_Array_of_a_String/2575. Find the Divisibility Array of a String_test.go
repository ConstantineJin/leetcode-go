package main

import (
	"reflect"
	"testing"
)

func Test_divisibilityArray(t *testing.T) {
	type args struct {
		word string
		m    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				word: "998244353",
				m:    3,
			},
			want: []int{1, 1, 0, 0, 0, 1, 1, 0, 0},
		},
		{
			name: "Example2",
			args: args{
				word: "1010",
				m:    10,
			},
			want: []int{0, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisibilityArray(tt.args.word, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divisibilityArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
