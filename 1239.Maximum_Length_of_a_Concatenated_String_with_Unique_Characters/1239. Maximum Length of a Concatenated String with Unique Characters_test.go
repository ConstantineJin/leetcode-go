package main

import "testing"

func Test_maxLength(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{arr: []string{"un", "iq", "ue"}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{arr: []string{"cha", "r", "act", "ers"}},
			want: 6,
		},
		{
			name: "Example3",
			args: args{arr: []string{"abcdefghijklmnopqrstuvwxyz"}},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.arr); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
