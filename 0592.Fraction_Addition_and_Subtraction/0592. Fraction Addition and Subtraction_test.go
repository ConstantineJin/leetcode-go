package main

import "testing"

func Test_fractionAddition(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{expression: "-1/2+1/2"},
			want: "0/1",
		},
		{
			name: "Example2",
			args: args{expression: "-1/2+1/2+1/3"},
			want: "1/3",
		},
		{
			name: "Example3",
			args: args{expression: "1/3-1/2"},
			want: "-1/6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionAddition(tt.args.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
