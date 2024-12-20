package main

import "testing"

func Test_minAnagramLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{s: "abba"},
			want: 2,
		},
		{
			name: "Example2",
			args: args{s: "cdef"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAnagramLength(tt.args.s); got != tt.want {
				t.Errorf("minAnagramLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
