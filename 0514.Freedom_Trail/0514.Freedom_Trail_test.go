package main

import "testing"

func Test_findRotateSteps(t *testing.T) {
	type args struct {
		ring string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				ring: "godding",
				key:  "gd",
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				ring: "godding",
				key:  "godding",
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findRotateSteps(tt.args.ring, tt.args.key); gotAns != tt.want {
				t.Errorf("findRotateSteps() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
