package main

import "testing"

func Test_numberOfAlternatingGroups(t *testing.T) {
	type args struct {
		colors []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{colors: []int{1, 1, 1}},
			wantAns: 0,
		},
		{
			name:    "Example2",
			args:    args{colors: []int{0, 1, 0, 0, 1}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfAlternatingGroups(tt.args.colors); gotAns != tt.wantAns {
				t.Errorf("numberOfAlternatingGroups() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
