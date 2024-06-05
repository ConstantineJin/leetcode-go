package main

import "testing"

func Test_minimumSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{s: "101"},
			want: 1,
		},
		{
			name: "Example2",
			args: args{s: "100"},
			want: 2,
		},
		{
			name: "Example3",
			args: args{s: "0111"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumSteps(tt.args.s); gotAns != tt.want {
				t.Errorf("minimumSteps() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
