package main

import "testing"

func Test_maximumInvitations(t *testing.T) {
	type args struct {
		favorite []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{favorite: []int{2, 2, 1, 2}},
			want: 3,
		},
		{
			name: "Example2",
			args: args{favorite: []int{1, 2, 0}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{favorite: []int{3, 0, 1, 4, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumInvitations(tt.args.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}
