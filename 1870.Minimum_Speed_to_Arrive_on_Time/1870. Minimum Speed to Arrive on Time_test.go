package main

import "testing"

func Test_minSpeedOnTime(t *testing.T) {
	type args struct {
		dist []int
		hour float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				dist: []int{1, 3, 2},
				hour: 6,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				dist: []int{1, 3, 2},
				hour: 2.7,
			},
			want: 3,
		},
		{
			name: "Example3",
			args: args{
				dist: []int{1, 3, 2},
				hour: 1.9,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
