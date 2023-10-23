package main

import "testing"

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "Example1",
			args:    args{satisfaction: []int{-1, -8, 0, 5, -9}},
			wantRes: 14,
		},
		{
			name:    "Example2",
			args:    args{satisfaction: []int{4, 3, 2}},
			wantRes: 20,
		},
		{
			name:    "Example3",
			args:    args{satisfaction: []int{-1, -4, -5}},
			wantRes: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxSatisfaction(tt.args.satisfaction); gotRes != tt.wantRes {
				t.Errorf("maxSatisfaction() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
