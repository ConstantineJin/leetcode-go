package main

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{strs: []string{"aba", "cdc", "eae"}},
			want: 3,
		},
		{
			name: "Example2",
			args: args{strs: []string{"aaa", "aaa", "aa"}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.strs); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
