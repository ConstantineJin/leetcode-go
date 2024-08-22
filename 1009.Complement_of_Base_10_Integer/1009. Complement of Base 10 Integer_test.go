package main

import "testing"

func Test_bitwiseComplement(t *testing.T) {
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
			want: 2,
		},
		{
			name: "Example2",
			args: args{n: 7},
			want: 0,
		},
		{
			name: "Example3",
			args: args{n: 10},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.n); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
