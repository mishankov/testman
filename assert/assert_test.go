package assert_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mishankov/testman/assert"
)

func TestAsserts(t *testing.T) {
	buffer := &bytes.Buffer{}
	ft := &FakeTesting{Buffer: buffer}

	testCases := []struct {
		name       string
		got        func() bool
		wantResult bool
		wantBuffer string
	}{
		{"True passing", func() bool { return assert.True(ft, true) }, true, ""},
		{"True not passing", func() bool { return assert.True(ft, false) }, false, "condition expected to be true"},
		{"Equal passing", func() bool { return assert.Equal(ft, 1, 1) }, true, ""},
		{"Equal not passing", func() bool { return assert.Equal(ft, 1, 2) }, false, "got 1 want 2"},
		{"DeepEqual passing", func() bool { return assert.DeepEqual(ft, []int{1, 2}, []int{1, 2}) }, true, ""},
		{"DeepEqual not passing", func() bool { return assert.DeepEqual(ft, []int{1, 2}, []int{2, 2}) }, false, "got [1 2] want [2 2]"},
		{"Contains passing", func() bool { return assert.Contains(ft, "some string", "me st") }, true, ""},
		{"Contains not passing", func() bool { return assert.Contains(ft, "some string", "hello") }, false, `expected "some string" to contain "hello"`},
		{"Regex passing", func() bool { return assert.Regex(ft, "111", `\d{3}`) }, true, ""},
		{"Regex not passing", func() bool { return assert.Regex(ft, "aaa", `\d{3}`) }, false, `"aaa" didn't match regexp "\d{3}"`},
		{"Regex not compiling", func() bool { return assert.Regex(ft, "aaa", `\p`) }, false, "regexp \"\\p\" didn't compile: error parsing regexp: invalid character class range: `\\p`"},
	}

	for _, test := range testCases {
		ft.Buffer.Reset()
		t.Run(test.name, func(t *testing.T) {
			if got := test.got(); got != test.wantResult {
				t.Errorf("got %v, want %v", got, test.wantResult)
			}

			if ft.Buffer.String() != test.wantBuffer {
				t.Errorf("got %q, want %q", ft.Buffer.String(), test.wantBuffer)
			}
		})
	}
}

type FakeTesting struct {
	Buffer *bytes.Buffer
}

func (ft *FakeTesting) Error(args ...interface{}) {
	ft.Buffer.WriteString(fmt.Sprint(args...))
}

func (ft *FakeTesting) Errorf(format string, args ...interface{}) {
	ft.Buffer.WriteString(fmt.Sprintf(format, args...))
}

func (ft *FakeTesting) Helper() {}
