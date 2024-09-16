package main

import "testing"

func Test_findTheLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{s: "eleetminicoworoep"},
			wantAns: 13,
		},
		{
			name:    "Example2",
			args:    args{s: "leetcodeisgreat"},
			wantAns: 5,
		},
		{
			name:    "Example3",
			args:    args{s: "bcbcbc"},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findTheLongestSubstring(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("findTheLongestSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
