package collector

import "testing"

func TestParseStringToBool(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},  // missing field: must not read as enabled
		{" ", false}, // whitespace only
		{"0", false},
		{"1", true},
		{"false", false},
		{"False", false},
		{"true", true},
		{"no", false},
		{"yes", true},
		{"off", false},
		{"on", true},
		{"disabled", false},
		{"enabled", true},
		{"2", true}, // unknown non-empty values keep the historical behaviour
	}
	for _, c := range cases {
		if got := parseStringToBool(c.in); got != c.want {
			t.Errorf("parseStringToBool(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
