package main

import "testing"

func reverseStringHelper(s []byte) (ans []byte) {
	reverseString(s)
	return s
}

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Example1",
			args: args{s: []byte{'h', 'e', 'l', 'l', 'o'}},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "Example2",
			args: args{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseStringHelper(tt.args.s)
		})
	}
}
