package main

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			wantAns: 8,
		},
		{
			name: "Example2",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     0,
			},
			wantAns: 6,
		},
		{
			name: "Example3",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				n:     2,
			},
			wantAns: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := leastInterval(tt.args.tasks, tt.args.n); gotAns != tt.wantAns {
				t.Errorf("leastInterval() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
