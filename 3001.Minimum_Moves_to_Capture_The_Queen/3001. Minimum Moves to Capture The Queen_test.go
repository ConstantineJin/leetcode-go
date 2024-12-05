package main

import "testing"

func Test_minMovesToCaptureTheQueen(t *testing.T) {
	type args struct{ a, b, c, d, e, f int }
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				a: 1,
				b: 1,
				c: 8,
				d: 8,
				e: 2,
				f: 3,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				a: 5,
				b: 3,
				c: 3,
				d: 4,
				e: 5,
				f: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMovesToCaptureTheQueen(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e, tt.args.f); got != tt.want {
				t.Errorf("minMovesToCaptureTheQueen() = %v, want %v", got, tt.want)
			}
		})
	}
}
