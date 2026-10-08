package greet

import "testing"

// TestHello checks Hello for normal, empty, whitespace-only and padded names.
//
// ENG: Each case runs as a subtest, e.g. TestHello/empty.
//
// LAO: ກວດ Hello ຫຼາຍກໍລະນີ — ແຕ່ລະກໍລະນີແມ່ນ subtest.
func TestHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "normal name", input: "Ton", want: "Hello, Ton!"},
		{name: "empty", input: "", want: "Hello, stranger!"},
		{name: "whitespace only", input: "   ", want: "Hello, stranger!"},
		{name: "padded name", input: "  Ton  ", want: "Hello, Ton!"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Hello(tc.input)
			if got != tc.want {
				t.Errorf("Hello(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
