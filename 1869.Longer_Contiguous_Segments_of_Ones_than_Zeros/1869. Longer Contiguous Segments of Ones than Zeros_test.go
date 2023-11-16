package main

import "testing"

func Test_checkZeroOnes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{s: "1101"},
			want: true,
		},
		{
			name: "Example2",
			args: args{s: "111000"},
			want: false,
		},
		{
			name: "Example3",
			args: args{s: "110100010"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkZeroOnes(tt.args.s); got != tt.want {
				t.Errorf("checkZeroOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
