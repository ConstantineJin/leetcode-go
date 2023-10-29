package main

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name  string
		args  args
		wantH int
	}{
		{
			name:  "Example1",
			args:  args{citations: []int{3, 0, 6, 1, 5}},
			wantH: 3,
		},
		{
			name:  "Example2",
			args:  args{citations: []int{1, 3, 1}},
			wantH: 1,
		},
		{
			name:  "Example3",
			args:  args{citations: []int{0, 1}},
			wantH: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotH := hIndex(tt.args.citations); gotH != tt.wantH {
				t.Errorf("hIndex() = %v, want %v", gotH, tt.wantH)
			}
		})
	}
}
