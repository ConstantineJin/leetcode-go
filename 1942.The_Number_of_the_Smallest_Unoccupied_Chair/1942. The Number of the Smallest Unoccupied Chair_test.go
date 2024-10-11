package main

import "testing"

func Test_smallestChair(t *testing.T) {
	type args struct {
		times        [][]int
		targetFriend int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				times:        [][]int{{1, 4}, {2, 3}, {4, 6}},
				targetFriend: 1,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				times:        [][]int{{3, 10}, {1, 5}, {2, 6}},
				targetFriend: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestChair(tt.args.times, tt.args.targetFriend); got != tt.want {
				t.Errorf("smallestChair() = %v, want %v", got, tt.want)
			}
		})
	}
}
