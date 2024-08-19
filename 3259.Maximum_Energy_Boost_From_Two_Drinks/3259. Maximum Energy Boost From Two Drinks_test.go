package main

import "testing"

func Test_maxEnergyBoost(t *testing.T) {
	type args struct {
		energyDrinkA []int
		energyDrinkB []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				energyDrinkA: []int{1, 3, 1},
				energyDrinkB: []int{3, 1, 1},
			},
			want: 5,
		},
		{
			name: "Example2",
			args: args{
				energyDrinkA: []int{4, 1, 1},
				energyDrinkB: []int{1, 1, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnergyBoost(tt.args.energyDrinkA, tt.args.energyDrinkB); got != tt.want {
				t.Errorf("maxEnergyBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}
