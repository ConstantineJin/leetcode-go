package main

import "testing"

func Test_addMinimum(t *testing.T) {
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
			args:    args{word: "b"},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{word: "aaa"},
			wantAns: 6,
		},
		{
			name:    "Example3",
			args:    args{word: "abc"},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := addMinimum(tt.args.word); gotAns != tt.wantAns {
				t.Errorf("addMinimum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
