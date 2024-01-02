package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{strs: []string{"flower", "flow", "flight"}},
			wantAns: "fl",
		},
		{
			name:    "Example2",
			args:    args{strs: []string{"dog", "racecar", "car"}},
			wantAns: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestCommonPrefix(tt.args.strs); gotAns != tt.wantAns {
				t.Errorf("longestCommonPrefix() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
