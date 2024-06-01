package main

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		n     int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:     5,
				limit: 2,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				n:     3,
				limit: 3,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.n, tt.args.limit); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
