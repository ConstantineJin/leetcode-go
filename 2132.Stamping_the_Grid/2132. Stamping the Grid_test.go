package main

import "testing"

func Test_possibleToStamp(t *testing.T) {
	type args struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				grid:        [][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}},
				stampHeight: 4,
				stampWidth:  3,
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				grid:        [][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
				stampHeight: 2,
				stampWidth:  2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleToStamp(tt.args.grid, tt.args.stampHeight, tt.args.stampWidth); got != tt.want {
				t.Errorf("possibleToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
