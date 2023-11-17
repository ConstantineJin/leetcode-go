package main

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				people: []int{1, 2},
				limit:  3,
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				people: []int{3, 2, 2, 1},
				limit:  3,
			},
			wantAns: 3,
		},
		{
			name: "Example3",
			args: args{
				people: []int{3, 5, 3, 4},
				limit:  5,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numRescueBoats(tt.args.people, tt.args.limit); gotAns != tt.wantAns {
				t.Errorf("numRescueBoats() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
