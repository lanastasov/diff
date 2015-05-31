// +build gofuzz

package diff

func Fuzz(data []byte) int {
	median := len(data) / 2
	a := string(data[:median])
	b := ""
	if median+1 < len(data) {
		b = string(data[median+1:])
	}
	ColouredDiff(a, b, true)
	return 1
}
