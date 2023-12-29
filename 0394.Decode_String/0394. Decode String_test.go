package main

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name:    "Example1",
			args:    args{s: "3[a]2[bc]"},
			wantRes: "aaabcbc",
		},
		{
			name:    "Example2",
			args:    args{s: "3[a2[c]]"},
			wantRes: "accaccacc",
		},
		{
			name:    "Example3",
			args:    args{s: "2[abc]3[cd]ef"},
			wantRes: "abcabccdcdcdef",
		},
		{
			name:    "Example4",
			args:    args{s: "abc3[cd]xyz"},
			wantRes: "abccdcdcdxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := decodeString(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("decodeString() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
