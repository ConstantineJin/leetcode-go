package main

import "testing"

func Test_numFriendRequests(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{ages: []int{16, 16}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{ages: []int{16, 17, 18}},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{ages: []int{20, 30, 100, 110, 120}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numFriendRequests(tt.args.ages); gotAns != tt.wantAns {
				t.Errorf("numFriendRequests() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
