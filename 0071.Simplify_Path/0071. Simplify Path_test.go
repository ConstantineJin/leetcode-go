package main

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{path: "/home/"},
			wantAns: "/home",
		},
		{
			name:    "Example2",
			args:    args{path: "/../"},
			wantAns: "/",
		},
		{
			name:    "Example3",
			args:    args{path: "/home//foo/"},
			wantAns: "/home/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := simplifyPath(tt.args.path); gotAns != tt.wantAns {
				t.Errorf("simplifyPath() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
