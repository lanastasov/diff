package diff

import (
	"testing"
)

func TestColor(t *testing.T) {
	got := color("string to color", 41)
	const want = "\u001b[41mstring to color\u001b[0m"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestColorLines(t *testing.T) {
	got := colorLines("str1\nstr2", 42)
	const want = "\u001b[42mstr1\u001b[0m\n\u001b[42mstr2\u001b[0m"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestColouredDiff(t *testing.T) {
	got := ColouredDiff(`abc def ghi
jkl mno
stu
yz
`, `abc ghi
jkl mno pqr
vwx
yz
`, false)
	const want = "abc \u001b[41mdef \u001b[0mghi\n" +
		"jkl mno\u001b[42m pqr\u001b[0m\n" +
		"\u001b[42mvwx\u001b[0m\u001b[41mstu\u001b[0m\n" +
		"yz\n"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
