package main

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k          int64
		operations []int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example1",
			args: args{
				k:          5,
				operations: []int{0, 0, 0},
			},
			want: 'a',
		},
		{
			name: "Example2",
			args: args{
				k:          10,
				operations: []int{0, 1, 0, 1},
			},
			want: 'b',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k, tt.args.operations); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
