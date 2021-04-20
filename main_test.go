package main

import (
	"testing"
)

func Test_noName(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not the same length", args{a: "abc", b: "abcd"}, false},
		{"same length but have no same letter", args{a: "cat", b: "dog"}, false},
		{"same length and have one same letter", args{a: "cat", b: "cog"}, false},
		{"all the same", args{a: "cat", b: "cat"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noName(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("no_name() = %v, want %v", got, tt.want)
			}
		})
	}
}
