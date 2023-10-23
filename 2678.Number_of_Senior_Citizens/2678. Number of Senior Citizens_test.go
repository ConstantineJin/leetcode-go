package main

import "testing"

func Test_countSeniors(t *testing.T) {
	type args struct {
		details []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "Example1",
			args:    args{details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}},
			wantRes: 2,
		},
		{
			name:    "Example2",
			args:    args{details: []string{"1313579440F2036", "2921522980M5644"}},
			wantRes: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countSeniors(tt.args.details); gotRes != tt.wantRes {
				t.Errorf("countSeniors() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
