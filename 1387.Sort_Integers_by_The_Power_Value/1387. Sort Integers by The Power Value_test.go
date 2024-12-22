package main

import "testing"

func Test_getKth(t *testing.T) {
	type args struct {
		lo int
		hi int
		k  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				lo: 12,
				hi: 15,
				k:  2,
			},
			want: 13,
		},
		{
			name: "Example2",
			args: args{
				lo: 7,
				hi: 11,
				k:  4,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKth(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
