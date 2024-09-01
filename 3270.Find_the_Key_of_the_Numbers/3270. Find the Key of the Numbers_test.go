package main

import "testing"

func Test_generateKey(t *testing.T) {
	type args struct {
		num1 int
		num2 int
		num3 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				num1: 1,
				num2: 10,
				num3: 1000,
			},
			want: 0,
		},
		{
			name: "Example2",
			args: args{
				num1: 987,
				num2: 879,
				num3: 798,
			},
			want: 777,
		},
		{
			name: "Example3",
			args: args{
				num1: 1,
				num2: 2,
				num3: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateKey(tt.args.num1, tt.args.num2, tt.args.num3); got != tt.want {
				t.Errorf("generateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
