package main

import "testing"

func Test_entityParser(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{text: "&amp; is an HTML entity but &ambassador; is not."},
			want: "& is an HTML entity but &ambassador; is not.",
		},
		{
			name: "Example2",
			args: args{text: "and I quote: &quot;...&quot;"},
			want: "and I quote: \"...\"",
		},
		{
			name: "Example3",
			args: args{text: "Stay home! Practice on Leetcode :)"},
			want: "Stay home! Practice on Leetcode :)",
		},
		{
			name: "Example4",
			args: args{text: "x &gt; y &amp;&amp; x &lt; y is always false"},
			want: "x > y && x < y is always false",
		},
		{
			name: "Example5",
			args: args{text: "leetcode.com&frasl;problemset&frasl;all"},
			want: "leetcode.com/problemset/all",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entityParser(tt.args.text); got != tt.want {
				t.Errorf("entityParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
