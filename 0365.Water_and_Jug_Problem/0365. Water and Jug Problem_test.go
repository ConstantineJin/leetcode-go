package main

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		jug1Capacity   int
		jug2Capacity   int
		targetCapacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				jug1Capacity:   3,
				jug2Capacity:   5,
				targetCapacity: 4,
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				jug1Capacity:   2,
				jug2Capacity:   6,
				targetCapacity: 5,
			},
			want: false,
		},
		{
			name: "Example3",
			args: args{
				jug1Capacity:   1,
				jug2Capacity:   2,
				targetCapacity: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.jug1Capacity, tt.args.jug2Capacity, tt.args.targetCapacity); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
