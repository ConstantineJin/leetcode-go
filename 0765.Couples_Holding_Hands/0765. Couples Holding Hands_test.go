package main

import "testing"

func Test_minSwapsCouples(t *testing.T) {
	type args struct {
		row []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{row: []int{0, 2, 1, 3}},
			want: 1,
		},
		{
			name: "Example2",
			args: args{row: []int{3, 2, 0, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwapsCouples(tt.args.row); got != tt.want {
				t.Errorf("minSwapsCouples() = %v, want %v", got, tt.want)
			}
		})
	}
}
