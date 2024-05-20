package main

import (
	"reflect"
	"testing"
)

func Test_canMakePaliQueries(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example1",
			args: args{
				s:       "abcda",
				queries: [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}},
			},
			want: []bool{true, false, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakePaliQueries(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMakePaliQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
