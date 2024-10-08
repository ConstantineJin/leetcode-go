package main

import "testing"

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				employees: []*Employee{{Id: 1, Importance: 5, Subordinates: []int{2, 3}}, {Id: 2, Importance: 3}, {Id: 3, Importance: 3}},
				id:        1,
			},
			want: 11,
		},
		{
			name: "Example2",
			args: args{
				employees: []*Employee{{Id: 1, Importance: 2, Subordinates: []int{5}}, {Id: 5, Importance: -3}},
				id:        5,
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}
