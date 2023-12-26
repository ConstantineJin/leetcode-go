package main

import "testing"

func Test_maxStudents(t *testing.T) {
	type args struct {
		seats [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{seats: [][]byte{{'#', '.', '#', '#', '.', '#'}, {'.', '#', '#', '#', '#', '.'}, {'#', '.', '#', '#', '.', '#'}}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{seats: [][]byte{{'.', '#'}, {'#', '#'}, {'#', '.'}, {'#', '#'}, {'.', '#'}}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{seats: [][]byte{{'#', '.', '.', '.', '#'}, {'.', '#', '.', '#', '.'}, {'.', '.', '#', '.', '.'}, {'#', '.', '.', '.', '#'}}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxStudents(tt.args.seats); got != tt.want {
				t.Errorf("maxStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
