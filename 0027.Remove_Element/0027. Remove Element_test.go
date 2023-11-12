package main

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{3, 2, 2, 3}, val: 3},
			wantCnt: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			wantCnt: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := removeElement(tt.args.nums, tt.args.val); gotCnt != tt.wantCnt {
				t.Errorf("removeElement() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
