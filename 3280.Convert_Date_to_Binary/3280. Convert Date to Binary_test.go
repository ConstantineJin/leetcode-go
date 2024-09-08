package main

import "testing"

func Test_convertDateToBinary(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{date: "2080-02-29"},
			want: "100000100000-10-11101",
		},
		{
			name: "Example2",
			args: args{date: "1900-01-01"},
			want: "11101101100-1-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDateToBinary(tt.args.date); got != tt.want {
				t.Errorf("convertDateToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
