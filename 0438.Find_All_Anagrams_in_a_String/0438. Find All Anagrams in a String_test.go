package main

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name:    "Example1",
			args:    args{s: "cbaebabacd", p: "abc"},
			wantRes: []int{0, 6},
		},
		{
			name:    "Example2",
			args:    args{s: "abab", p: "ab"},
			wantRes: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findAnagrams() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
