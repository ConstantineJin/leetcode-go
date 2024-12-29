package main

import "testing"

func Test_rankTeams(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"}},
			want: "ACB",
		},
		{
			name: "Example2",
			args: args{votes: []string{"WXYZ", "XYZW"}},
			want: "XWYZ",
		},
		{
			name: "Example3",
			args: args{votes: []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}},
			want: "ZMNAGUEDSJYLBOPHRQICWFXTVK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankTeams(tt.args.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
