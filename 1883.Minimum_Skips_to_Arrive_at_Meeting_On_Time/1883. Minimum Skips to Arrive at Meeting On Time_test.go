package main

import "testing"

func Test_minSkips(t *testing.T) {
	type args struct {
		dist        []int
		speed       int
		hoursBefore int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				dist:        []int{1, 3, 2},
				speed:       4,
				hoursBefore: 2,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				dist:        []int{7, 3, 5, 5},
				speed:       2,
				hoursBefore: 10,
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				dist:        []int{7, 3, 5, 5},
				speed:       1,
				hoursBefore: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSkips(tt.args.dist, tt.args.speed, tt.args.hoursBefore); got != tt.want {
				t.Errorf("minSkips() = %v, want %v", got, tt.want)
			}
		})
	}
}
