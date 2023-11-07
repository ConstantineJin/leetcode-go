package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "Example1",
			args:    args{s: "abcabcbb"},
			wantRes: 3,
		},
		{
			name:    "Example2",
			args:    args{s: "bbbbb"},
			wantRes: 1,
		},
		{
			name:    "Example3",
			args:    args{s: "pwwkew"},
			wantRes: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := lengthOfLongestSubstring(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
