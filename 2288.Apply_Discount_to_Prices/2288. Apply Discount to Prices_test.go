package main

import "testing"

func Test_discountPrices(t *testing.T) {
	type args struct {
		sentence string
		discount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				sentence: "there are $1 $2 and 5$ candies in the shop",
				discount: 50,
			},
			want: "there are $0.50 $1.00 and 5$ candies in the shop",
		},
		{
			name: "Example2",
			args: args{
				sentence: "1 2 $3 4 $5 $6 7 8$ $9 $10$",
				discount: 100,
			},
			want: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := discountPrices(tt.args.sentence, tt.args.discount); got != tt.want {
				t.Errorf("discountPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
