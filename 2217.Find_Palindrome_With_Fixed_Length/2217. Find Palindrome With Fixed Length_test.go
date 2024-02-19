package main

import (
	"reflect"
	"testing"
)

func Test_kthPalindrome(t *testing.T) {
	type args struct {
		queries   []int
		intLength int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Example1",
			args: args{
				queries:   []int{1, 2, 3, 4, 5, 90},
				intLength: 3,
			},
			want: []int64{101, 111, 121, 131, 141, 999},
		},
		{
			name: "Example2",
			args: args{
				queries:   []int{2, 4, 6},
				intLength: 4,
			},
			want: []int64{1111, 1331, 1551},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthPalindrome(tt.args.queries, tt.args.intLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
