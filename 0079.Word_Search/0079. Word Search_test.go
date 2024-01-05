package main

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name    string
		args    args
		wantAns bool
	}{
		{
			name: "Example1",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCCED",
			},
			wantAns: true,
		},
		{
			name: "Example2",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "SEE",
			},
			wantAns: true,
		},
		{
			name: "Example3",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCB",
			},
			wantAns: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := exist(tt.args.board, tt.args.word); gotAns != tt.wantAns {
				t.Errorf("exist() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
