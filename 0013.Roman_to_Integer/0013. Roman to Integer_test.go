package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{s: "III"},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{s: "IV"},
			wantAns: 4,
		},
		{
			name:    "Example3",
			args:    args{s: "IX"},
			wantAns: 9,
		},
		{
			name:    "Example4",
			args:    args{s: "LVIII"},
			wantAns: 58,
		},
		{
			name:    "Example5",
			args:    args{s: "MCMXCIV"},
			wantAns: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := romanToInt(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("romanToInt() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
