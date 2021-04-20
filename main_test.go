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
		{"all empty", args{a: "", b: ""}, true},
		{"same letters but with different order", args{a: "cat", b: "tac"}, true},
		{"recurring letters", args{a: "aaa", b: "aaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noName(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("no_name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"dummy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_utilityFunction(t *testing.T) {
	type args struct {
		s string
		j int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{"abc", 0}, "bc"},
		{"simple 1", args{"abc", 1}, "ac"},
		{"simple 1", args{"abc", 2}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utilityFunction(tt.args.s, tt.args.j); got != tt.want {
				t.Errorf("utilityFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
