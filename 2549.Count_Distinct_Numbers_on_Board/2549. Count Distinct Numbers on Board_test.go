package main

import "testing"

func Test_distinctIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{n: 5},
			want: 4,
		},
		{
			name: "Example2",
			args: args{n: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctIntegers(tt.args.n); got != tt.want {
				t.Errorf("distinctIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
