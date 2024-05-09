package main

import "testing"

func Test_minimumRefill(t *testing.T) {
	type args struct {
		plants    []int
		capacityA int
		capacityB int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				plants:    []int{2, 2, 3, 3},
				capacityA: 5,
				capacityB: 5,
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				plants:    []int{2, 2, 3, 3},
				capacityA: 3,
				capacityB: 4,
			},
			wantAns: 2,
		},
		{
			name: "Example3",
			args: args{
				plants:    []int{5},
				capacityA: 10,
				capacityB: 8,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumRefill(tt.args.plants, tt.args.capacityA, tt.args.capacityB); gotAns != tt.wantAns {
				t.Errorf("minimumRefill() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
