package main

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		source   string
		target   string
		original []byte
		changed  []byte
		cost     []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				source:   "abcd",
				target:   "acbe",
				original: []byte{'a', 'b', 'c', 'c', 'e', 'd'},
				changed:  []byte{'b', 'c', 'b', 'e', 'b', 'e'},
				cost:     []int{2, 5, 5, 1, 2, 20},
			},
			wantAns: 28,
		},
		{
			name: "Example2",
			args: args{
				source:   "aaaa",
				target:   "bbbb",
				original: []byte{'a', 'c'},
				changed:  []byte{'c', 'b'},
				cost:     []int{1, 2},
			},
			wantAns: 12,
		},
		{
			name: "Example3",
			args: args{
				source:   "abcd",
				target:   "abce",
				original: []byte{'a'},
				changed:  []byte{'e'},
				cost:     []int{10000},
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumCost(tt.args.source, tt.args.target, tt.args.original, tt.args.changed, tt.args.cost); gotAns != tt.wantAns {
				t.Errorf("minimumCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
