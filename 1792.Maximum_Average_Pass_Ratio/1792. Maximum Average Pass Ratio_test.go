package main

import "testing"

func Test_maxAverageRatio(t *testing.T) {
	type args struct {
		classes       [][]int
		extraStudents int
	}
	tests := []struct {
		name    string
		args    args
		wantAns float64
	}{
		{
			name: "Example1",
			args: args{
				classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
				extraStudents: 2,
			},
			wantAns: 47.0 / 60,
		},
		{
			name: "Example2",
			args: args{
				classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
				extraStudents: 4,
			},
			wantAns: 353.0 / 660,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxAverageRatio(tt.args.classes, tt.args.extraStudents); gotAns != tt.wantAns {
				t.Errorf("maxAverageRatio() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
