package main

import "testing"

func Test_maximumBags(t *testing.T) {
	type args struct {
		capacity        []int
		rocks           []int
		additionalRocks int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				capacity:        []int{2, 3, 4, 5},
				rocks:           []int{1, 2, 4, 4},
				additionalRocks: 2,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				capacity:        []int{10, 2, 2},
				rocks:           []int{2, 2, 0},
				additionalRocks: 100,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumBags(tt.args.capacity, tt.args.rocks, tt.args.additionalRocks); gotAns != tt.wantAns {
				t.Errorf("maximumBags() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
