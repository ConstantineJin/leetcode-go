package main

import (
	"reflect"
	"testing"
)

func Test_numOfBurgers(t *testing.T) {
	type args struct {
		tomatoSlices int
		cheeseSlices int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				tomatoSlices: 16,
				cheeseSlices: 7,
			},
			want: []int{1, 6},
		},
		{
			name: "Example2",
			args: args{
				tomatoSlices: 17,
				cheeseSlices: 4,
			},
			want: []int{},
		},
		{
			name: "Example3",
			args: args{
				tomatoSlices: 4,
				cheeseSlices: 17,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfBurgers(tt.args.tomatoSlices, tt.args.cheeseSlices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numOfBurgers() = %v, want %v", got, tt.want)
			}
		})
	}
}
