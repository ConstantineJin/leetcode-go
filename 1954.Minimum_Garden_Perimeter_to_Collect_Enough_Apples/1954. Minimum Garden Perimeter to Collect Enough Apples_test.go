package main

import "testing"

func Test_minimumPerimeter(t *testing.T) {
	type args struct {
		neededApples int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{neededApples: 1},
			want: 8,
		},
		{
			name: "Example2",
			args: args{neededApples: 13},
			want: 16,
		},
		{
			name: "Example3",
			args: args{neededApples: 1000000000},
			want: 5040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPerimeter(tt.args.neededApples); got != tt.want {
				t.Errorf("minimumPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
