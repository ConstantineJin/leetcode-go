package main

import "testing"

func Test_theMaximumAchievableX(t *testing.T) {
	type args struct {
		num int
		t   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				num: 4,
				t:   1,
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				num: 3,
				t:   2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theMaximumAchievableX(tt.args.num, tt.args.t); got != tt.want {
				t.Errorf("theMaximumAchievableX() = %v, want %v", got, tt.want)
			}
		})
	}
}
