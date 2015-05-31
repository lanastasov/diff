// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package diff

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func allWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func tokenize(s string) (tokens []string) {
	var buf bytes.Buffer
	previousIsLetter := false
	previousIsSpace := false
	for _, r := range s {
		isSpace := unicode.IsSpace(r)
		isLetter := !isSpace && unicode.IsLetter(r)
		if isLetter != previousIsLetter || isSpace != previousIsSpace {
			if buf.Len() > 0 {
				tokens = append(tokens, buf.String())
				buf.Reset()
			}
		}
		buf.WriteRune(r)
		previousIsLetter = isLetter
		previousIsSpace = isSpace
	}

	if buf.Len() > 0 {
		tokens = append(tokens, buf.String())
	}
	return
}

// NumberLines numbers lines in a string
func NumberLines(str string) string {
	lines := strings.Split(str, "\n")
	numpad := len(strconv.Itoa(len(lines)))
	for i, s := range lines {
		lines[i] = fmt.Sprintf("%*d | %s", numpad, i+1, s)
	}
	return strings.Join(lines, "\n")
}

// Indent indents lines in a string with n spaces.
func Indent(str string, n int) string {
	padding := strings.Repeat(" ", n)
	lines := strings.Split(str, "\n")
	for i, s := range lines {
		lines[i] = padding + s
	}
	return strings.Join(lines, "\n")
}
