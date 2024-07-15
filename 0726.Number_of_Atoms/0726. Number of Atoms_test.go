package main

import "testing"

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{s: "H2O"},
			wantAns: "H2O",
		},
		{
			name:    "Example2",
			args:    args{s: "Mg(OH)2"},
			wantAns: "H2MgO2",
		},
		{
			name:    "Example3",
			args:    args{s: "K4(ON(SO3)2)2"},
			wantAns: "K4N2O14S4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countOfAtoms(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("countOfAtoms() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
