package main

import "testing"

func Test_kthFactor(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n: 12,
				k: 3,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				n: 7,
				k: 2,
			},
			want: 7,
		},
		{
			name: "Example3",
			args: args{
				n: 4,
				k: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthFactor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
