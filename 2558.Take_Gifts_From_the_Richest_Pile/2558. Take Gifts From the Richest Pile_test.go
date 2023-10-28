package main

import "testing"

func Test_pickGifts(t *testing.T) {
	type args struct {
		gifts []int
		k     int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int64
	}{
		{
			name: "Example1",
			args: args{
				gifts: []int{25, 64, 9, 4, 100},
				k:     4,
			},
			wantSum: 29,
		},
		{
			name: "Example2",
			args: args{
				gifts: []int{1, 1, 1, 1},
				k:     4,
			},
			wantSum: 4,
		},
		{
			name: "Example3",
			args: args{
				gifts: []int{37, 46, 17, 40, 50, 54, 11, 1, 25, 43, 21, 31, 29, 58, 49, 73, 54, 5, 52, 73, 54, 6, 22, 58, 9, 34, 21, 58, 68, 63, 72, 1, 5, 64, 42, 40, 60, 7, 54, 25},
				k:     36,
			},
			wantSum: 220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := pickGifts(tt.args.gifts, tt.args.k); gotSum != tt.wantSum {
				t.Errorf("pickGifts() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
