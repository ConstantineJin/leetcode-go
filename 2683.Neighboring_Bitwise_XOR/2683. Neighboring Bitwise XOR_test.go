package main

import "testing"

func Test_doesValidArrayExist(t *testing.T) {
	type args struct {
		derived []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{derived: []int{1, 1, 0}},
			want: true,
		},
		{
			name: "Example2",
			args: args{derived: []int{1, 1}},
			want: true,
		},
		{
			name: "Example3",
			args: args{derived: []int{1, 0}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesValidArrayExist(tt.args.derived); got != tt.want {
				t.Errorf("doesValidArrayExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
