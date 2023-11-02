package main

import "testing"

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name:    "Example1",
			args:    args{"B0B6G0R6R0R6G9"},
			wantCnt: 1,
		},
		{
			name:    "Example2",
			args:    args{"B0R0G0R9R0B0G0"},
			wantCnt: 1,
		},
		{
			name:    "Example3",
			args:    args{"G4"},
			wantCnt: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := countPoints(tt.args.rings); gotCnt != tt.wantCnt {
				t.Errorf("countPoints() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
