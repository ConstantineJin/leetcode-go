package main

import "testing"

func Test_distanceTraveled(t *testing.T) {
	type args struct {
		mainTank       int
		additionalTank int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				mainTank:       5,
				additionalTank: 10,
			},
			want: 60,
		},
		{
			name: "Example2",
			args: args{
				mainTank:       1,
				additionalTank: 2,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := distanceTraveled(tt.args.mainTank, tt.args.additionalTank); gotAns != tt.want {
				t.Errorf("distanceTraveled() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
