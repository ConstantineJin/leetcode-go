package main

import "testing"

func Test_uniqueLetterString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "'Example1",
			args:    args{s: "ABC"},
			wantAns: 10,
		},
		{
			name:    "'Example2",
			args:    args{s: "ABA"},
			wantAns: 8,
		},
		{
			name:    "'Example3",
			args:    args{s: "LEETCODE"},
			wantAns: 92,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := uniqueLetterString(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("uniqueLetterString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
