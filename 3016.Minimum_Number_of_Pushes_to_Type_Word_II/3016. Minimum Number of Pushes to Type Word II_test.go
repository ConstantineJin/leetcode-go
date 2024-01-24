package main

import "testing"

func Test_minimumPushes(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{word: "abcde"},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{word: "xyzxyzxyzxyz"},
			wantAns: 12,
		},
		{
			name:    "Example3",
			args:    args{word: "aabbccddeeffgghhiiiiii"},
			wantAns: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumPushes(tt.args.word); gotAns != tt.wantAns {
				t.Errorf("minimumPushes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
