package main

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
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
			args:    args{words: []string{"bella", "label", "roller"}},
			wantAns: []string{"e", "l", "l"},
		}, {
			name:    "Example2",
			args:    args{words: []string{"cool", "lock", "cook"}},
			wantAns: []string{"c", "o"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := commonChars(tt.args.words); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("commonChars() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
