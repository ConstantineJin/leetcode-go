package main

import "testing"

func Test_maximumSubsequenceCount(t *testing.T) {
	type args struct {
		text    string
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				text:    "abdcdbc",
				pattern: "ac",
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				text:    "aabb",
				pattern: "ab",
			},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumSubsequenceCount(tt.args.text, tt.args.pattern); gotAns != tt.wantAns {
				t.Errorf("maximumSubsequenceCount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
