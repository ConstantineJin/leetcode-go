package main

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{n: 12},
			wantAns: 21,
		},
		{
			name:    "Example2",
			args:    args{n: 21},
			wantAns: -1,
		},
		{
			name:    "Example3",
			args:    args{n: 8341},
			wantAns: 8413,
		},
		{
			name:    "Example4",
			args:    args{n: 8243},
			wantAns: 8324,
		},
		{
			name:    "Example5",
			args:    args{n: 2147483476},
			wantAns: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := nextGreaterElement(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("nextGreaterElement() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
