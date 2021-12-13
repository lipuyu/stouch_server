package utils

import (
	"stouch_server/src/common/utils"
	"testing"
)

func TestUtils(t *testing.T) {
	cases := []struct {
		ok bool
		a, b, want string
	}{
		{true, "Hello a", "Hello b", "Hello a"},
		{false, "Hello a", "Hello b", "Hello b"},
	}
	for _, c := range cases {
		got := utils.If(c.ok, c.a, c.b)
		if got != c.want {
			t.Errorf("if(%t, %q, %q) == %q, want %q", c.ok, c.a, c.b, got, c.want)
		}
	}
}
