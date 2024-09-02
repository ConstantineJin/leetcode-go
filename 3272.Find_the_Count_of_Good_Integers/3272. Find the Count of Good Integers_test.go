package main

import "testing"

func Test_countGoodIntegers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				n: 3,
				k: 5,
			},
			wantAns: 27,
		},
		{
			name: "Example2",
			args: args{
				n: 1,
				k: 4,
			},
			wantAns: 2,
		},
		{
			name: "Example3",
			args: args{
				n: 5,
				k: 6,
			},
			wantAns: 2468,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countGoodIntegers(tt.args.n, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countGoodIntegers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
