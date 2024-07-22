package main

import "testing"

func Test_doesAliceWin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{s: "leetcoder"},
			want: true,
		},
		{
			name: "Example2",
			args: args{s: "bbcd"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesAliceWin(tt.args.s); got != tt.want {
				t.Errorf("doesAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
