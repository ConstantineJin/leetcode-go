package main

import "testing"

func Test_countBattleships(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{board: [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', '.'}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{board: [][]byte{{'.'}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countBattleships(tt.args.board); gotAns != tt.wantAns {
				t.Errorf("countBattleships() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
