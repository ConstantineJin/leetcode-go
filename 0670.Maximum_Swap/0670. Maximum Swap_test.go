package main

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{num: 2736},
			want: 7236,
		},
		{
			name: "Example2",
			args: args{num: 9973},
			want: 9973,
		},
		{
			name: "Example3",
			args: args{num: 7324},
			want: 7423,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
