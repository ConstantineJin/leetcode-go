package main

import "testing"

func Test_capitalizeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{title: "capiTalIze tHe titLe"},
			wantAns: "Capitalize The Title",
		},
		{
			name:    "Example2",
			args:    args{title: "First leTTeR of EACH Word"},
			wantAns: "First Letter of Each Word",
		},
		{
			name:    "Example3",
			args:    args{title: "i lOve leetcode"},
			wantAns: "i Love Leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := capitalizeTitle(tt.args.title); gotAns != tt.wantAns {
				t.Errorf("capitalizeTitle() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
