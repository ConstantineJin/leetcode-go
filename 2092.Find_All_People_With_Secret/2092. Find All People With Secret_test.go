package main

import (
	"reflect"
	"testing"
)

func Test_findAllPeople(t *testing.T) {
	type args struct {
		n           int
		meetings    [][]int
		firstPerson int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				n:           6,
				meetings:    [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}},
				firstPerson: 1,
			},
			wantAns: []int{0, 1, 2, 3, 5},
		},
		{
			name: "Example2",
			args: args{
				n:           4,
				meetings:    [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}},
				firstPerson: 3,
			},
			wantAns: []int{0, 1, 3},
		},
		{
			name: "Example3",
			args: args{
				n:           5,
				meetings:    [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}},
				firstPerson: 1,
			},
			wantAns: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAllPeople(tt.args.n, tt.args.meetings, tt.args.firstPerson); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAllPeople() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
