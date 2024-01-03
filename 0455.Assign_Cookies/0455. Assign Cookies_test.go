package main

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				g: []int{1, 2, 3},
				s: []int{1, 1},
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				g: []int{1, 2},
				s: []int{1, 2, 3},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findContentChildren(tt.args.g, tt.args.s); gotAns != tt.wantAns {
				t.Errorf("findContentChildren() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
