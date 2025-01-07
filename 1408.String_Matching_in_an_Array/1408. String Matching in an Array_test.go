package main

import (
	"reflect"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "Example1",
			args:    args{words: []string{"mass", "as", "hero", "superhero"}},
			wantAns: []string{"as", "hero"},
		},
		{
			name:    "Example2",
			args:    args{words: []string{"leetcode", "et", "code"}},
			wantAns: []string{"et", "code"},
		},
		{
			name:    "Example3",
			args:    args{words: []string{"blue", "green", "bu"}},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := stringMatching(tt.args.words); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("stringMatching() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
