package main

import (
	"reflect"
	"testing"
)

func equal(got [][]string, want [][]string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	for _, row := range got {
		for _, item := range row {
			map1[item]++
		}
	}
	for _, row := range want {
		for _, item := range row {
			map2[item]++
		}
	}
	return reflect.DeepEqual(map1, map2)
}

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]string
	}{
		{
			name:    "Example1",
			args:    args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			wantRes: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:    "Example2",
			args:    args{strs: []string{""}},
			wantRes: [][]string{{""}},
		},
		{
			name:    "Example3",
			args:    args{strs: []string{"a"}},
			wantRes: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := groupAnagrams(tt.args.strs); !equal(gotRes, tt.wantRes) {
				t.Errorf("groupAnagrams() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
