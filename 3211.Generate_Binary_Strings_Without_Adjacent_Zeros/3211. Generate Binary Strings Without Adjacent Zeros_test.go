package main

import (
	"reflect"
	"testing"
)

func Test_validStrings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "Example1",
			args:    args{n: 3},
			wantAns: []string{"010", "011", "101", "110", "111"},
		},
		{
			name:    "Example2",
			args:    args{n: 1},
			wantAns: []string{"0", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := validStrings(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("validStrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
