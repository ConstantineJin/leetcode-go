package main

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example1",
			args: args{k: 5},
			want: 'b',
		},
		{
			name: "Example2",
			args: args{k: 10},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
