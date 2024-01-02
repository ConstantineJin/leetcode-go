package main

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "Example1",
			args:    args{digits: "23"},
			wantAns: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:    "Example2",
			args:    args{digits: ""},
			wantAns: nil,
		},
		{
			name:    "Example3",
			args:    args{digits: "2"},
			wantAns: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := letterCombinations(tt.args.digits); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("letterCombinations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
