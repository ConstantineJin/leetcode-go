package main

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{equations: []string{"a==b", "b!=a"}},
			want: false,
		},
		{
			name: "Example2",
			args: args{equations: []string{"b==a", "a==b"}},
			want: true,
		},
		{
			name: "Example3",
			args: args{equations: []string{"a==b", "b==c", "a==c"}},
			want: true,
		},
		{
			name: "Example4",
			args: args{equations: []string{"a==b", "b!=c", "c==a"}},
			want: false,
		},
		{
			name: "Example5",
			args: args{equations: []string{"c==c", "b==d", "x!=z"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
