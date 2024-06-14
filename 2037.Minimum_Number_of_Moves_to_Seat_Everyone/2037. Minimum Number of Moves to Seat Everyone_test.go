package main

import "testing"

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				seats:    []int{3, 1, 5},
				students: []int{2, 7, 4},
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				seats:    []int{4, 1, 5, 9},
				students: []int{1, 3, 2, 6},
			},
			wantAns: 7,
		},
		{
			name: "Example3",
			args: args{
				seats:    []int{2, 2, 6, 6},
				students: []int{1, 3, 2, 6},
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minMovesToSeat(tt.args.seats, tt.args.students); gotAns != tt.wantAns {
				t.Errorf("minMovesToSeat() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
