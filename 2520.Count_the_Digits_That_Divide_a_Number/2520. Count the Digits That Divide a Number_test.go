package main

import "testing"

func Test_countDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name:    "Example1",
			args:    args{num: 7},
			wantCnt: 1,
		},
		{
			name:    "Example2",
			args:    args{num: 121},
			wantCnt: 2,
		},
		{
			name:    "Example3",
			args:    args{num: 1248},
			wantCnt: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := countDigits(tt.args.num); gotCnt != tt.wantCnt {
				t.Errorf("countDigits() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
