package main

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "Example1",
			args: args{
				num1: "11",
				num2: "123",
			},
			wantAns: "134",
		},
		{
			name: "Example2",
			args: args{
				num1: "456",
				num2: "77",
			},
			wantAns: "533",
		},
		{
			name: "Example3",
			args: args{
				num1: "0",
				num2: "0",
			},
			wantAns: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := addStrings(tt.args.num1, tt.args.num2); gotAns != tt.wantAns {
				t.Errorf("addStrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
