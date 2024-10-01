package main

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		in1    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Example1",
			args: args{
				target: []int{1, 3},
				in1:    3,
			},
			wantAns: []string{"Push", "Push", "Pop", "Push"},
		},
		{
			name: "Example2",
			args: args{
				target: []int{1, 2, 3},
				in1:    3,
			},
			wantAns: []string{"Push", "Push", "Push"},
		},
		{
			name: "Example3",
			args: args{
				target: []int{1, 2},
				in1:    4,
			},
			wantAns: []string{"Push", "Push"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := buildArray(tt.args.target, tt.args.in1); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("buildArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
