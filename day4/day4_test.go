package main

import "testing"

func Test_byr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{name: "test1", args: args{"2002"}, want: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byr(tt.args.s); got != tt.want {
				t.Errorf("byr() = %v, want %v", got, tt.want)
			}
		})
	}
}

