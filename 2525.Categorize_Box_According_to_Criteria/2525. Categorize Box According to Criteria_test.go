package main

import "testing"

func Test_categorizeBox(t *testing.T) {
	type args struct {
		length int
		width  int
		height int
		mass   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				length: 1000,
				width:  35,
				height: 700,
				mass:   300,
			},
			want: "Heavy",
		},
		{
			name: "Example2",
			args: args{
				length: 200,
				width:  50,
				height: 800,
				mass:   50,
			},
			want: "Neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := categorizeBox(tt.args.length, tt.args.width, tt.args.height, tt.args.mass); got != tt.want {
				t.Errorf("categorizeBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
