package main

import "testing"

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				deadends: []string{"0201", "0101", "0102", "1212", "2002"},
				target:   "0202",
			},
			wantAns: 6,
		},
		{
			name: "Example2",
			args: args{
				deadends: []string{"8888"},
				target:   "0009",
			},
			wantAns: 1,
		},
		{
			name: "Example3",
			args: args{
				deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
				target:   "8888",
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := openLock(tt.args.deadends, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("openLock() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
