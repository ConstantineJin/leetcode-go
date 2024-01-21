package main

import (
	"reflect"
	"testing"
)

func Test_splitWordsBySeparator(t *testing.T) {
	type args struct {
		words     []string
		separator byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Example1",
			args: args{
				words:     []string{"one.two.three", "four.five", "six"},
				separator: '.',
			},
			wantAns: []string{"one", "two", "three", "four", "five", "six"},
		},
		{
			name: "Example2",
			args: args{
				words:     []string{"$easy$", "$problem$"},
				separator: '$',
			},
			wantAns: []string{"easy", "problem"},
		},
		{
			name: "Example3",
			args: args{
				words:     []string{"|||"},
				separator: '|',
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := splitWordsBySeparator(tt.args.words, tt.args.separator); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("splitWordsBySeparator() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
