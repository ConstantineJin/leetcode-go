package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "Example1",
			args:    args{n: 3},
			wantAns: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:    "Example2",
			args:    args{n: 1},
			wantAns: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := generateParenthesis(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("generateParenthesis() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
