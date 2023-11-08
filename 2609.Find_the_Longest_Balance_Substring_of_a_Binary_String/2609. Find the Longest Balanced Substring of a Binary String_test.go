package main

import "testing"

func Test_findTheLongestBalancedSubstring(t *testing.T) {
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
			args:    args{s: "01000111"},
			wantRes: 6,
		},
		{
			name:    "Example2",
			args:    args{s: "00111"},
			wantRes: 4,
		},
		{
			name:    "Example3",
			args:    args{s: "111"},
			wantRes: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findTheLongestBalancedSubstring(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("findTheLongestBalancedSubstring() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
