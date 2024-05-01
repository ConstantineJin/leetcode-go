package main

import "testing"

func Test_numberOfEmployeesWhoMetTarget(t *testing.T) {
	type args struct {
		hours  []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				hours:  []int{0, 1, 2, 3, 4},
				target: 2,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				hours:  []int{5, 1, 4, 2, 2},
				target: 6,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfEmployeesWhoMetTarget(tt.args.hours, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("numberOfEmployeesWhoMetTarget() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
