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

// Package diff implements coloured diffs on a word per word basis.
package diff

const (
	diffNone = iota
	diffAdded
	diffRemoved
)

type diff struct {
	ignoreWhitespace bool
}

type component struct {
	value  string
	status int
}

type path struct {
	newPos int
	comp   []component
}

func clonePath(p *path) *path {
	newp := path{
		newPos: p.newPos,
		comp:   make([]component, len(p.comp)),
	}
	copy(newp.comp, p.comp)
	return &newp
}

func (d *diff) extractCommon(basePath *path, new, old []string, diagonalPath int) int {
	newPos := basePath.newPos
	oldPos := newPos - diagonalPath
	for newPos+1 < len(new) && oldPos+1 < len(old) && d.equals(new[newPos+1], old[oldPos+1]) {
		newPos++
		oldPos++
		d.pushComponent(&basePath.comp, new[newPos], diffNone)
	}
	basePath.newPos = newPos
	return oldPos
}

func (d *diff) equals(a, b string) bool {
	if d.ignoreWhitespace &&
		allWhitespace(a) &&
		allWhitespace(b) {
		return true
	}
	return a == b
}

func (d *diff) pushComponent(comp *[]component, value string, status int) {
	if len(*comp) > 0 {
		last := (*comp)[len(*comp)-1]
		if last.status == status {
			(*comp)[len(*comp)-1] = component{
				value:  last.value + value,
				status: status,
			}
			return
		}
	}

	*comp = append(*comp, component{
		value:  value,
		status: status,
	})
}

func (d *diff) diff(a, b string) []component {
	if a == b {
		return []component{{value: b}}
	}
	if b == "" {
		return []component{{value: a, status: diffRemoved}}
	}
	if a == "" {
		return []component{{value: b, status: diffAdded}}
	}

	old := tokenize(a)
	new := tokenize(b)

	maxEditLength := len(old) + len(new)
	bestPath := map[int]*path{
		0: &path{newPos: -1},
	}
	oldPos := d.extractCommon(bestPath[0], new, old, 0)
	if bestPath[0].newPos+1 >= len(new) && oldPos+1 >= len(old) {
		return bestPath[0].comp
	}

	for editLength := 1; editLength <= maxEditLength; editLength++ {
		for diagonalPath := -1 * editLength; diagonalPath <= editLength; diagonalPath += 2 {
			addPath := bestPath[diagonalPath-1]
			removePath := bestPath[diagonalPath+1]
			oldPos = 0
			if removePath != nil {
				oldPos = removePath.newPos
			}
			oldPos -= diagonalPath
			if addPath != nil {
				delete(bestPath, diagonalPath-1)
			}

			canAdd := addPath != nil && addPath.newPos+1 < len(new)
			canRemove := removePath != nil && 0 <= oldPos && oldPos < len(old)
			if !canAdd && !canRemove {
				delete(bestPath, diagonalPath)
				continue
			}

			var basePath *path
			if !canAdd || (canRemove && addPath.newPos < removePath.newPos) {
				basePath = clonePath(removePath)
				d.pushComponent(&basePath.comp, old[oldPos], diffRemoved)
			} else {
				basePath = clonePath(addPath)
				basePath.newPos++
				d.pushComponent(&basePath.comp, new[basePath.newPos], diffAdded)
			}

			oldPos = d.extractCommon(basePath, new, old, diagonalPath)

			if basePath.newPos+1 >= len(new) && oldPos+1 >= len(old) {
				return basePath.comp
			}

			bestPath[diagonalPath] = basePath
		}
	}
	return nil
}
