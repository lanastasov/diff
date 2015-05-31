package diff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type testCase struct {
		a, b string
		ws   bool
		want []component
	}
	testCases := []testCase{
		{"xxx", "xxx", false, []component{{value: "xxx"}}},
		{"xxx", "", false, []component{{value: "xxx", status: diffRemoved}}},
		{"", "xxx", false, []component{{value: "xxx", status: diffAdded}}},
		{"aaa bbb", "aaa    bbb", false, []component{{value: "aaa"}, {value: "    ", status: diffAdded}, {value: " ", status: diffRemoved}, {value: "bbb"}}},
		{"aaa bbb", "aaa    bbb", true, []component{{value: "aaa    bbb"}}},
	}
	for _, tc := range testCases {
		d := diff{tc.ws}
		got := d.diff(tc.a, tc.b)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("want %#v, got %#v", tc.want, got)
		}
	}
}
