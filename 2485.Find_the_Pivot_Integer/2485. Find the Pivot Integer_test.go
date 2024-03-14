package main

import "testing"

func Test_pivotInteger(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantX int
	}{
		{
			name:  "Example1",
			args:  args{n: 8},
			wantX: 6,
		},
		{
			name:  "Example2",
			args:  args{n: 1},
			wantX: 1,
		},
		{
			name:  "Example3",
			args:  args{n: 4},
			wantX: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotX := pivotInteger(tt.args.n); gotX != tt.wantX {
				t.Errorf("pivotInteger() = %v, want %v", gotX, tt.wantX)
			}
		})
	}
}
