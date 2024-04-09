package main

import "testing"

func Test_firstDayBeenInAllRooms(t *testing.T) {
	type args struct {
		nextVisit []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{nextVisit: []int{0, 0}},
			want: 2,
		},
		{
			name: "Example2",
			args: args{nextVisit: []int{0, 0, 2}},
			want: 6,
		},
		{
			name: "Example3",
			args: args{nextVisit: []int{0, 1, 2, 0}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := firstDayBeenInAllRooms(tt.args.nextVisit); gotAns != tt.want {
				t.Errorf("firstDayBeenInAllRooms() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
