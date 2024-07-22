package main

import "testing"

func Test_minChanges(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				n: 13,
				k: 4,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				n: 21,
				k: 21,
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				n: 14,
				k: 13,
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minChanges(tt.args.n, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("minChanges() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
