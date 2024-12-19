package main

import (
	"reflect"
	"testing"
)

func Test_stableMountains(t *testing.T) {
	type args struct {
		height    []int
		threshold int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				height:    []int{1, 2, 3, 4, 5},
				threshold: 2,
			},
			wantAns: []int{3, 4},
		},
		{
			name: "Example2",
			args: args{
				height:    []int{10, 1, 10, 1, 10},
				threshold: 3,
			},
			wantAns: []int{1, 3},
		},
		{
			name: "Example3",
			args: args{
				height:    []int{10, 1, 10, 1, 10},
				threshold: 10,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := stableMountains(tt.args.height, tt.args.threshold); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("stableMountains() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
