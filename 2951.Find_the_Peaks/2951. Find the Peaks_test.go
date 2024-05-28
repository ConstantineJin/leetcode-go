package main

import (
	"reflect"
	"testing"
)

func Test_findPeaks(t *testing.T) {
	type args struct {
		mountain []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name:    "Example1",
			args:    args{mountain: []int{2, 4, 4}},
			wantAns: nil,
		},
		{
			name:    "Example2",
			args:    args{mountain: []int{1, 4, 3, 8, 5}},
			wantAns: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findPeaks(tt.args.mountain); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findPeaks() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
