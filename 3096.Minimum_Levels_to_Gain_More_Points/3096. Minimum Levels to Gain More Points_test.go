package main

import "testing"

func Test_minimumLevels(t *testing.T) {
	type args struct {
		possible []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{possible: []int{1, 0, 1, 0}},
			want: 1,
		},
		{
			name: "Example2",
			args: args{possible: []int{1, 1, 1, 1, 1}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{possible: []int{0, 0}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLevels(tt.args.possible); got != tt.want {
				t.Errorf("minimumLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
