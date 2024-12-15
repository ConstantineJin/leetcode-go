package main

import "testing"

func Test_minSetSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{arr: []int{7, 7, 7, 7, 7, 7}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minSetSize(tt.args.arr); gotAns != tt.wantAns {
				t.Errorf("minSetSize() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
