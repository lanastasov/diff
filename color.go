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
	"fmt"
	"strings"
)

const (
	diffAddedColor   = 42
	diffRemovedColor = 41
)

func color(s string, clr int) string {
	return fmt.Sprintf("\u001b[%dm%s\u001b[0m", clr, s)
}

func colorLines(str string, clr int) string {
	lines := strings.Split(str, "\n")
	for i, s := range lines {
		lines[i] = color(s, clr)
	}
	return strings.Join(lines, "\n")
}

// ColouredDiff compares two strings and returns a coloured difference between them.
func ColouredDiff(a, b string, ignoreWhitespace bool) string {
	d := diff{ignoreWhitespace}
	comp := d.diff(a, b)
	lines := make([]string, 0, len(comp))
	for _, c := range comp {
		switch c.status {
		case diffAdded:
			lines = append(lines, colorLines(c.value, diffAddedColor))
		case diffRemoved:
			lines = append(lines, colorLines(c.value, diffRemovedColor))
		default:
			lines = append(lines, c.value)
		}
	}
	return strings.Join(lines, "")
}
