package main

import "testing"

func Test_minDamage(t *testing.T) {
	type args struct {
		power  int
		damage []int
		health []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				power:  4,
				damage: []int{1, 2, 3, 4},
				health: []int{4, 5, 6, 8},
			},
			wantAns: 39,
		},
		{
			name: "Example2",
			args: args{
				power:  1,
				damage: []int{1, 1, 1, 1},
				health: []int{1, 2, 3, 4},
			},
			wantAns: 20,
		},
		{
			name: "Example3",
			args: args{
				power:  8,
				damage: []int{40},
				health: []int{59},
			},
			wantAns: 320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDamage(tt.args.power, tt.args.damage, tt.args.health); gotAns != tt.wantAns {
				t.Errorf("minDamage() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
