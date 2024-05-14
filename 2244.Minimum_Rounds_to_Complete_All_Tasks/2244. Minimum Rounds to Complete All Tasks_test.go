package main

import "testing"

func Test_minimumRounds(t *testing.T) {
	type args struct {
		tasks []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{tasks: []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{tasks: []int{2, 2, 3}},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumRounds(tt.args.tasks); gotAns != tt.wantAns {
				t.Errorf("minimumRounds() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
