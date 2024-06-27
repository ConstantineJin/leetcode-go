package main

import "testing"

func Test_smallestString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{str: "cbabc"},
			want: "baabc",
		},
		{
			name: "Example2",
			args: args{str: "acbbc"},
			want: "abaab",
		},
		{
			name: "Example3",
			args: args{str: "leetcode"},
			want: "kddsbncd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestString(tt.args.str); got != tt.want {
				t.Errorf("smallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
