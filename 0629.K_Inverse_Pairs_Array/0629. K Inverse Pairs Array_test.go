package main

import "testing"

func Test_kInversePairs(t *testing.T) {
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
				n: 3,
				k: 0,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				n: 3,
				k: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kInversePairs(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kInversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
