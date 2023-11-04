package main

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{
			name:    "Example1",
			args:    args{s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"},
			wantRes: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name:    "Example2",
			args:    args{s: "AAAAAAAAAAAAA"},
			wantRes: []string{"AAAAAAAAAA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findRepeatedDnaSequences(tt.args.s); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
