package diff

import (
	"reflect"
	"testing"
)

func TestAllWhitespace(t *testing.T) {
	type testCase struct {
		in   string
		want bool
	}
	testCases := []testCase{
		{"", true},
		{" ", true},
		{"a", false},
		{" a b c ", false},
		{"\u00a0", true},
		{" \u00a0 ", true},
	}
	for _, tc := range testCases {
		got := allWhitespace(tc.in)
		if got != tc.want {
			t.Errorf("allWhitespace(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestTokenize(t *testing.T) {
	type testCase struct {
		in   string
		want []string
	}
	testCases := []testCase{
		{"", nil},
		{"a", []string{"a"}},
		{"abc", []string{"abc"}},
		{" ", []string{" "}},
		{"        ", []string{"        "}},
		{" abc", []string{" ", "abc"}},
		{"abc ", []string{"abc", " "}},
		{"!", []string{"!"}},
		{"!?!", []string{"!?!"}},
		{"<em>", []string{"<", "em", ">"}},
		{" ?", []string{" ", "?"}},
		{"? ", []string{"?", " "}},
		{"abc!123", []string{"abc", "!123"}},
		{"aaa        bbb", []string{"aaa", "        ", "bbb"}},
		{" \u00a0 ", []string{" \u00a0 "}},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
			[]string{
				"Lorem", " ", "ipsum", " ", "dolor", " ", "sit", " ", "amet", ",", " ",
				"consectetur", " ", "adipiscing", " ", "elit",
			},
		},
		{
			"Иногда простым трёхчасовым ацетилированием простой домашней сметаны в фтороводородной среде с палладиевым катализом можно добиться удивительных результатов.",
			[]string{
				"Иногда", " ", "простым", " ", "трёхчасовым", " ", "ацетилированием", " ",
				"простой", " ", "домашней", " ", "сметаны", " ", "в", " ", "фтороводородной",
				" ", "среде", " ", "с", " ", "палладиевым", " ", "катализом", " ", "можно", " ",
				"добиться", " ", "удивительных", " ", "результатов", ".",
			},
		},
	}
	for i, tc := range testCases {
		got := tokenize(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("#%d: want %#v, got %#v", i+1, tc.want, got)
		}
	}
}

func TestIndent(t *testing.T) {
	type testCase struct {
		in     string
		indent int
		want   string
	}
	testCases := []testCase{
		{"", 5, "     "},
		{"abc", 0, "abc"},
		{"abc", 1, " abc"},
		{"abc", 8, "        abc"},
		{"abc\ndef\nghi", 4, "    abc\n    def\n    ghi"},
		{" abc\n\n  def", 4, "     abc\n    \n      def"},
	}
	for _, tc := range testCases {
		got := Indent(tc.in, tc.indent)
		if got != tc.want {
			t.Errorf("Indent(%q, %d) = %q, want %q", tc.in, tc.indent, got, tc.want)
		}
	}
}

func TestNumberLines(t *testing.T) {
	type testCase struct {
		in   string
		want string
	}
	testCases := []testCase{
		{"", "1 | "},
		{"abc", "1 | abc"},
		{"abc\ndef\nghi", "1 | abc\n2 | def\n3 | ghi"},
		{"\n\n\n", "1 | \n2 | \n3 | \n4 | "},
		{"a\nb\nc\nd\ne\nf\ng\nh\ni\nj", " 1 | a\n 2 | b\n 3 | c\n 4 | d\n 5 | e\n 6 | f\n 7 | g\n 8 | h\n 9 | i\n10 | j"},
	}
	for _, tc := range testCases {
		got := NumberLines(tc.in)
		if got != tc.want {
			t.Errorf("NumberLines(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
