package main

import "testing"

func Test_numberOfWeeks(t *testing.T) {
	type args struct {
		milestones []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{milestones: []int{1, 2, 3}},
			want: 6,
		},
		{
			name: "Example2",
			args: args{milestones: []int{5, 2, 1}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfWeeks(tt.args.milestones); gotAns != tt.want {
				t.Errorf("numberOfWeeks() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
