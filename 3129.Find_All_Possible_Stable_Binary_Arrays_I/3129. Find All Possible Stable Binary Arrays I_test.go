package main

import "testing"

func Test_numberOfStableArrays(t *testing.T) {
	type args struct {
		zero  int
		one   int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				zero:  1,
				one:   1,
				limit: 2,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				zero:  1,
				one:   2,
				limit: 1,
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				zero:  3,
				one:   3,
				limit: 2,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfStableArrays(tt.args.zero, tt.args.one, tt.args.limit); got != tt.want {
				t.Errorf("numberOfStableArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
