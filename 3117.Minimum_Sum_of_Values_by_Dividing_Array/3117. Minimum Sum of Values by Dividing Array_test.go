package main

import "testing"

func Test_minimumValueSum(t *testing.T) {
	type args struct {
		nums      []int
		andValues []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums:      []int{1, 4, 3, 3, 2},
				andValues: []int{0, 3, 3, 2},
			},
			want: 12,
		},
		{
			name: "Example2",
			args: args{
				nums:      []int{2, 3, 5, 7, 7, 7, 5},
				andValues: []int{0, 7, 5},
			},
			want: 17,
		},
		{
			name: "Example3",
			args: args{
				nums:      []int{1, 2, 3, 4},
				andValues: []int{2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumValueSum(tt.args.nums, tt.args.andValues); got != tt.want {
				t.Errorf("minimumValueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
