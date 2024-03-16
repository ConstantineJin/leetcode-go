package main

import "testing"

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name          string
		args          args
		wantMaxLength int
	}{
		{
			name:          "Example1",
			args:          args{nums: []int{0, 1}},
			wantMaxLength: 2,
		},
		{
			name:          "Example2",
			args:          args{nums: []int{0, 1, 0}},
			wantMaxLength: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxLength := findMaxLength(tt.args.nums); gotMaxLength != tt.wantMaxLength {
				t.Errorf("findMaxLength() = %v, want %v", gotMaxLength, tt.wantMaxLength)
			}
		})
	}
}
