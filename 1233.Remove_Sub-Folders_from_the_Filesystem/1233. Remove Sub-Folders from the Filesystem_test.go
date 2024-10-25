package main

import (
	"reflect"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	type args struct {
		folder []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "Example1",
			args:    args{folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}},
			wantAns: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name:    "Example2",
			args:    args{folder: []string{"/a", "/a/b/c", "/a/b/d"}},
			wantAns: []string{"/a"},
		},
		{
			name:    "Example3",
			args:    args{folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
			wantAns: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := removeSubfolders(tt.args.folder); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("removeSubfolders() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
