package main

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{num: 123},
			want: "One Hundred Twenty Three",
		},
		{
			name: "Example2",
			args: args{num: 12345},
			want: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name: "Example3",
			args: args{num: 1234567},
			want: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
