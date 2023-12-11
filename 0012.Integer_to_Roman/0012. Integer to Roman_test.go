package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{num: 3},
			wantAns: "III",
		},
		{
			name:    "Example2",
			args:    args{num: 4},
			wantAns: "IV",
		},
		{
			name:    "Example3",
			args:    args{num: 9},
			wantAns: "IX",
		},
		{
			name:    "Example4",
			args:    args{num: 58},
			wantAns: "LVIII",
		},
		{
			name:    "Example5",
			args:    args{num: 1994},
			wantAns: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := intToRoman(tt.args.num); gotAns != tt.wantAns {
				t.Errorf("intToRoman() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
