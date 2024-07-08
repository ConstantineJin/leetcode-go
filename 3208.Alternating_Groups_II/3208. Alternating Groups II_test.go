package main

import "testing"

func Test_numberOfAlternatingGroups(t *testing.T) {
	type args struct {
		colors []int
		k      int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				colors: []int{0, 1, 0, 1, 0},
				k:      3,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				colors: []int{0, 1, 0, 0, 1, 0, 1},
				k:      6,
			},
			wantAns: 2,
		},
		{
			name: "Example3",
			args: args{
				colors: []int{1, 1, 0, 1},
				k:      4,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfAlternatingGroups(tt.args.colors, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("numberOfAlternatingGroups() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
