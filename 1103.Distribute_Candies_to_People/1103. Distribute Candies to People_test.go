package main

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies   int
		numPeople int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				candies:   7,
				numPeople: 4,
			},
			want: []int{1, 2, 3, 1},
		},
		{
			name: "Example2",
			args: args{
				candies:   10,
				numPeople: 3,
			},
			want: []int{5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies, tt.args.numPeople); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
