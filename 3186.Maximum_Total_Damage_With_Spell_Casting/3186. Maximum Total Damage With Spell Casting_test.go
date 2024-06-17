package main

import "testing"

func Test_maximumTotalDamage(t *testing.T) {
	type args struct {
		power []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{power: []int{1, 1, 3, 4}},
			want: 6,
		},
		{
			name: "Example2",
			args: args{power: []int{7, 1, 6, 6}},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalDamage(tt.args.power); got != tt.want {
				t.Errorf("maximumTotalDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
