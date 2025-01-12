package main

import "testing"

func Test_canBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				s:      "))()))",
				locked: "010100",
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				s:      "()()",
				locked: "0000",
			},
			want: true,
		},
		{
			name: "Example3",
			args: args{
				s:      ")",
				locked: "0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
