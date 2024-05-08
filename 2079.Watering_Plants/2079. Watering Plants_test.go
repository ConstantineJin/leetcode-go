package main

import "testing"

func Test_wateringPlants(t *testing.T) {
	type args struct {
		plants   []int
		capacity int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				plants:   []int{2, 2, 3, 3},
				capacity: 5,
			},
			wantAns: 14,
		},
		{
			name: "Example2",
			args: args{
				plants:   []int{1, 1, 1, 4, 2, 3},
				capacity: 4,
			},
			wantAns: 30,
		},
		{
			name: "Example3",
			args: args{
				plants:   []int{7, 7, 7, 7, 7, 7, 7},
				capacity: 8,
			},
			wantAns: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := wateringPlants(tt.args.plants, tt.args.capacity); gotAns != tt.wantAns {
				t.Errorf("wateringPlants() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
