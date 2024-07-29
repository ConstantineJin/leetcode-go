package main

import "testing"

func Test_numTeams(t *testing.T) {
	type args struct {
		rating []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{rating: []int{2, 5, 3, 4, 1}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{rating: []int{2, 1, 3}},
			wantAns: 0,
		},
		{
			name:    "Example3",
			args:    args{rating: []int{1, 2, 3, 4}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numTeams(tt.args.rating); gotAns != tt.wantAns {
				t.Errorf("numTeams() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
