package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{logs: []string{"d1/", "d2/", "../", "d21/", "./"}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}},
			wantAns: 3,
		},
		{
			name:    "Example3",
			args:    args{logs: []string{"d1/", "../", "../", "../"}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minOperations(tt.args.logs); gotAns != tt.wantAns {
				t.Errorf("minOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
