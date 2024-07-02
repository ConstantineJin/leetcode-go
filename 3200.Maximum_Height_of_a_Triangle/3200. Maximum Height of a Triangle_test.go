package main

import "testing"

func Test_maxHeightOfTriangle(t *testing.T) {
	type args struct {
		red  int
		blue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				red:  2,
				blue: 4,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				red:  2,
				blue: 1,
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				red:  1,
				blue: 1,
			},
			want: 1,
		},
		{
			name: "Example4",
			args: args{
				red:  10,
				blue: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxHeightOfTriangle(tt.args.red, tt.args.blue); got != tt.want {
				t.Errorf("maxHeightOfTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
