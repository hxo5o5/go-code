package test

import (
	"gotest/gotesting/reverse"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello,world", "dlrow,olleh"},
		{"hello,世界", "界世,olleh"},
		{"", ""},
	}
	for _, c := range cases {
		got := reverse.Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q)=%q,want %q", c.in, got, c.want)
		}
	}
}
