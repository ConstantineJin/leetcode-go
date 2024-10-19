package main

import "testing"

func Test_findKthBit(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example1",
			args: args{
				n: 3,
				k: 1,
			},
			want: '0',
		},
		{
			name: "Example2",
			args: args{
				n: 4,
				k: 11,
			},
			want: '1',
		},
		{
			name: "Example3",
			args: args{
				n: 1,
				k: 1,
			},
			want: '0',
		},
		{
			name: "Example4",
			args: args{
				n: 2,
				k: 3,
			},
			want: '1',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthBit(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
