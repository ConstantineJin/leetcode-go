package main

import "testing"

func Test_minNonZeroProduct(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{p: 1},
			want: 1,
		},
		{
			name: "Example2",
			args: args{p: 2},
			want: 6,
		},
		{
			name: "Example3",
			args: args{p: 3},
			want: 1512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNonZeroProduct(tt.args.p); got != tt.want {
				t.Errorf("minNonZeroProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
