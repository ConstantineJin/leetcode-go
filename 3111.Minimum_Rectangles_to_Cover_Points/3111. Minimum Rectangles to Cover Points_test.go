package main

import "testing"

func Test_minRectanglesToCoverPoints(t *testing.T) {
	type args struct {
		points [][]int
		w      int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				points: [][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}},
				w:      1,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				points: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
				w:      2,
			},
			wantAns: 3,
		},
		{
			name: "Example3",
			args: args{
				points: [][]int{{2, 3}, {1, 2}},
				w:      0,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minRectanglesToCoverPoints(tt.args.points, tt.args.w); gotAns != tt.wantAns {
				t.Errorf("minRectanglesToCoverPoints() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
