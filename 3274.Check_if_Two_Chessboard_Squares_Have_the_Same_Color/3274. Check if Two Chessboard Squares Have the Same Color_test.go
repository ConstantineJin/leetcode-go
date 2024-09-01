package main

import "testing"

func Test_checkTwoChessboards(t *testing.T) {
	type args struct {
		coordinate1 string
		coordinate2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				coordinate1: "a1",
				coordinate2: "c3",
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				coordinate1: "a1",
				coordinate2: "h3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTwoChessboards(tt.args.coordinate1, tt.args.coordinate2); got != tt.want {
				t.Errorf("checkTwoChessboards() = %v, want %v", got, tt.want)
			}
		})
	}
}
