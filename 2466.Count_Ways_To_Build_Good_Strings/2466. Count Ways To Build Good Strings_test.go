package main

import "testing"

func Test_countGoodStrings(t *testing.T) {
	type args struct {
		low  int
		high int
		zero int
		one  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				low:  3,
				high: 3,
				zero: 1,
				one:  1,
			},
			wantAns: 8,
		},
		{
			name: "Example1",
			args: args{
				low:  2,
				high: 3,
				zero: 1,
				one:  2,
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countGoodStrings(tt.args.low, tt.args.high, tt.args.zero, tt.args.one); gotAns != tt.wantAns {
				t.Errorf("countGoodStrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
