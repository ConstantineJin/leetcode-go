package main

import "testing"

func Test_smallestBeautifulString(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				str: "abcz",
				k:   26,
			},
			want: "abda",
		},
		{
			name: "Example2",
			args: args{
				str: "dc",
				k:   4,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestBeautifulString(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("smallestBeautifulString() = %v, want %v", got, tt.want)
			}
		})
	}
}
