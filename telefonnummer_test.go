package telefonnummer

import (
	"testing"
)

func TestParseVoiceMail(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"222", "Röstbrevlåda"},
		{"333", "Röstbrevlåda"},
		{"888", "Röstbrevlåda"},
	}

	for _, c := range cases {
		got := Parse(c.in)
		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestParseLandline(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"4681234567", "08-123 45 67"},
		{"081234567", "08-123 45 67"},
		{"468123456", "08-12 34 56"},
		{"08123456", "08-12 34 56"},
		{"46812345678", "08-123 456 78"},
	}

	for _, c := range cases {
		got := Parse(c.in)
		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestParseMobile(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"46721834008", "072-183 40 08"},
		{"0721834008", "072-183 40 08"},
		{"0701234567", "070-123 45 67"},
	}

	for _, c := range cases {
		got := Parse(c.in)
		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestParseThreeDigitAreaCode(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"03112345", "031-123 45"},
		{"031446572", "031-44 65 72"},
		{"4631626262", "031-62 62 62"},
		{"031626262", "031-62 62 62"},
		{"46311234567", "031-123 45 67"},
		{"0311234567", "031-123 45 67"},
	}

	for _, c := range cases {
		got := Parse(c.in)

		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestParseFourDigitAreaCode(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"050012345", "0500-123 45"},
		{"46304123456", "0304-12 34 56"},
		{"0304123456", "0304-12 34 56"},
		{"46500123456", "0500-12 34 56"},
		{"0500123456", "0500-12 34 56"},
	}

	for _, c := range cases {
		got := Parse(c.in)

		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestParseInternational(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"+46701234567", "070-123 45 67"},
		{"+46 (0) 701234567", "070-123 45 67"},
	}

	for _, c := range cases {
		got := Parse(c.in)

		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
