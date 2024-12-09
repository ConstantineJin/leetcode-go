package main

import "testing"

func Test_squareIsWhite(t *testing.T) {
	type args struct {
		coordinates string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{coordinates: "a1"},
			want: false,
		},
		{
			name: "Example2",
			args: args{coordinates: "h3"},
			want: true,
		},
		{
			name: "Example3",
			args: args{coordinates: "c7"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareIsWhite(tt.args.coordinates); got != tt.want {
				t.Errorf("squareIsWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}
