package main

import "testing"

func Test_numberOfSubmatrices(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{grid: [][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]byte{{'X', 'X'}, {'X', 'Y'}}},
			wantAns: 0,
		},
		{
			name:    "Example3",
			args:    args{grid: [][]byte{{'.', '.'}, {'.', '.'}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfSubmatrices(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("numberOfSubmatrices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
