package main

import "testing"

func Test_maximumPopulation(t *testing.T) {
	type args struct {
		logs [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{logs: [][]int{{1993, 1999}, {2000, 2010}}},
			wantAns: 1993,
		},
		{
			name:    "Example2",
			args:    args{logs: [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}},
			wantAns: 1960,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumPopulation(tt.args.logs); gotAns != tt.wantAns {
				t.Errorf("maximumPopulation() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
