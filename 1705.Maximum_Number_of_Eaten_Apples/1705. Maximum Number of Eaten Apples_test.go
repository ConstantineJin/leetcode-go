package main

import "testing"

func Test_eatenApples(t *testing.T) {
	type args struct {
		apples []int
		days   []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				apples: []int{1, 2, 3, 5, 2},
				days:   []int{3, 2, 1, 4, 2},
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				apples: []int{3, 0, 0, 0, 0, 2},
				days:   []int{3, 0, 0, 0, 0, 2},
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := eatenApples(tt.args.apples, tt.args.days); gotAns != tt.wantAns {
				t.Errorf("eatenApples() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
