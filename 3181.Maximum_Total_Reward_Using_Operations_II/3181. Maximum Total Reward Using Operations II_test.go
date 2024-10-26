package main

import "testing"

func Test_maxTotalReward(t *testing.T) {
	type args struct {
		rewardValues []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{rewardValues: []int{1, 1, 3, 3}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{rewardValues: []int{1, 6, 4, 3, 2}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalReward(tt.args.rewardValues); got != tt.want {
				t.Errorf("maxTotalReward() = %v, want %v", got, tt.want)
			}
		})
	}
}
