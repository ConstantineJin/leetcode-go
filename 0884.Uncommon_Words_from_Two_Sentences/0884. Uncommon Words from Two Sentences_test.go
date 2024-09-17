package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Example1",
			args: args{
				s1: "this apple is sweet",
				s2: "this apple is sour",
			},
			wantAns: []string{"sour", "sweet"},
		},
		{
			name: "Example2",
			args: args{
				s1: "apple apple",
				s2: "banana",
			},
			wantAns: []string{"banana"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := uncommonFromSentences(tt.args.s1, tt.args.s2)
			sort.Strings(gotAns)
			sort.Strings(tt.wantAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("uncommonFromSentences() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
