package main

import "testing"

func Test_getSmallestString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				s: "zbbz",
				k: 3,
			},
			want: "aaaz",
		},
		{
			name: "Example2",
			args: args{
				s: "xaxcd",
				k: 4,
			},
			want: "aawcd",
		},
		{
			name: "Example3",
			args: args{
				s: "lol",
				k: 0,
			},
			want: "lol",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
