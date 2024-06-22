package main

import "testing"

func Test_maximumTripletValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{nums: []int{12, 6, 1, 2, 7}},
			want: 77,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 10, 3, 4, 19}},
			want: 133,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 2, 3}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTripletValue(tt.args.nums); got != tt.want {
				t.Errorf("maximumTripletValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
