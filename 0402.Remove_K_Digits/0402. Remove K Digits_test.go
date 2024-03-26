package main

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "Example1",
			args: args{
				num: "1432219",
				k:   3,
			},
			wantAns: "1219",
		},
		{
			name: "Example2",
			args: args{
				num: "10200",
				k:   1,
			},
			wantAns: "200",
		},
		{
			name: "Example3",
			args: args{
				num: "10",
				k:   2,
			},
			wantAns: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := removeKdigits(tt.args.num, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("removeKdigits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
