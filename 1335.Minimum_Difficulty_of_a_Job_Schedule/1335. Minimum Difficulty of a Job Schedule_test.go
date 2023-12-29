package main

import "testing"

func Test_minDifficulty(t *testing.T) {
	type args struct {
		jobDifficulty []int
		d             int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				jobDifficulty: []int{6, 5, 4, 3, 2, 1},
				d:             2,
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				jobDifficulty: []int{9, 9, 9},
				d:             4,
			},
			want: -1,
		},
		{
			name: "Example3",
			args: args{
				jobDifficulty: []int{1, 1, 1},
				d:             3,
			},
			want: 3,
		},
		{
			name: "Example4",
			args: args{
				jobDifficulty: []int{7, 1, 7, 1, 7, 1},
				d:             3,
			},
			want: 15,
		},
		{
			name: "Example5",
			args: args{
				jobDifficulty: []int{11, 111, 22, 222, 33, 333, 44, 444},
				d:             6,
			},
			want: 843,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficulty(tt.args.jobDifficulty, tt.args.d); got != tt.want {
				t.Errorf("minDifficulty() = %v, want %v", got, tt.want)
			}
		})
	}
}
