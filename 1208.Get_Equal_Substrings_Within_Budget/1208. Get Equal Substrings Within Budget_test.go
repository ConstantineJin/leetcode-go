package main

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				s:       "abcd",
				t:       "bcdf",
				maxCost: 3,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				s:       "abcd",
				t:       "cdef",
				maxCost: 3,
			},
			wantAns: 1,
		},
		{
			name: "Example3",
			args: args{
				s:       "abcd",
				t:       "acde",
				maxCost: 0,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); gotAns != tt.wantAns {
				t.Errorf("equalSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
