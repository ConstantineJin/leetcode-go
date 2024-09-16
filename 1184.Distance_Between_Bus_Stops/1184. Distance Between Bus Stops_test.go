package main

import "testing"

func Test_distanceBetweenBusStops(t *testing.T) {
	type args struct {
		distance    []int
		start       int
		destination int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 1,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 2,
			},
			want: 3,
		},
		{
			name: "Example3",
			args: args{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetweenBusStops(tt.args.distance, tt.args.start, tt.args.destination); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
