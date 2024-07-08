package main

import "testing"

func Test_maximumPoints(t *testing.T) {
	type args struct {
		enemyEnergies []int
		currentEnergy int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				enemyEnergies: []int{3, 2, 2},
				currentEnergy: 2,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				enemyEnergies: []int{2},
				currentEnergy: 10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPoints(tt.args.enemyEnergies, tt.args.currentEnergy); got != tt.want {
				t.Errorf("maximumPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
