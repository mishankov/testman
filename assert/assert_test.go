package assert_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mishankov/testman/assert"
)

func TestTrue(t *testing.T) {
	buffer := &bytes.Buffer{}
	ft := FakeTesting{Buffer: buffer}

	t.Run("passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.True(ft, true)

		if !res {
			t.Error("assert expected to pass")
		}

		if !(ft.Buffer.Len() == 0) {
			t.Error("buffer expected to be empty, got:", ft.Buffer.String())
		}
	})

	t.Run("not passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.True(ft, false)

		if res {
			t.Error("assert expected not to pass")
		}

		if ft.Buffer.String() != "condition expected to be true" {
			t.Error("buffer expected to have error message, got:", ft.Buffer.String())
		}
	})
}

func TestEqual(t *testing.T) {
	buffer := &bytes.Buffer{}
	ft := FakeTesting{Buffer: buffer}

	t.Run("passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Equal(ft, 1, 1)

		if !res {
			t.Error("assert expected to pass")
		}

		if !(ft.Buffer.Len() == 0) {
			t.Error("buffer expected to be empty, got:", ft.Buffer.String())
		}
	})

	t.Run("not passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Equal(ft, 1, 2)

		if res {
			t.Error("assert expected not to pass")
		}

		if ft.Buffer.String() != "got 1 want 2" {
			t.Error("buffer expected to have error message, got:", ft.Buffer.String())
		}
	})
}

func TestDeepEqual(t *testing.T) {
	buffer := &bytes.Buffer{}
	ft := FakeTesting{Buffer: buffer}

	t.Run("passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.DeepEqual(ft, []int{1, 2}, []int{1, 2})

		if !res {
			t.Error("assert expected to pass")
		}

		if !(ft.Buffer.Len() == 0) {
			t.Error("buffer expected to be empty, got:", ft.Buffer.String())
		}
	})

	t.Run("not passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.DeepEqual(ft, []int{1, 2}, []int{2, 2})

		if res {
			t.Error("assert expected not to pass")
		}

		if ft.Buffer.String() != "got [1 2] want [2 2]" {
			t.Error("buffer expected to have error message, got:", ft.Buffer.String())
		}
	})
}

func TestContains(t *testing.T) {
	buffer := &bytes.Buffer{}
	ft := FakeTesting{Buffer: buffer}

	t.Run("passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Contains(ft, "some string", "me st")

		if !res {
			t.Error("assert expected to pass")
		}

		if !(ft.Buffer.Len() == 0) {
			t.Error("buffer expected to be empty, got:", ft.Buffer.String())
		}
	})

	t.Run("not passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Contains(ft, "some string", "hello")

		if res {
			t.Error("assert expected not to pass")
		}

		if ft.Buffer.String() != `expected "some string" to contain "hello"` {
			t.Error("buffer expected to have error message, got:", ft.Buffer.String())
		}
	})
}

func TestRegex(t *testing.T) {
	buffer := &bytes.Buffer{}
	ft := FakeTesting{Buffer: buffer}

	t.Run("passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Regex(ft, "111", `\d{3}`)

		if !res {
			t.Error("assert expected to pass")
		}

		if !(ft.Buffer.Len() == 0) {
			t.Error("buffer expected to be empty, got:", ft.Buffer.String())
		}
	})

	t.Run("not passing", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Regex(ft, "aaa", `\d{3}`)

		if res {
			t.Error("assert expected not to pass")
		}

		if ft.Buffer.String() != `"aaa" didn't match regexp "\d{3}"` {
			t.Error("buffer expected to have error message, got:", ft.Buffer.String())
		}
	})

	t.Run("breaking", func(t *testing.T) {
		defer ft.Buffer.Reset()
		res := assert.Regex(ft, "aaa", `\p`)

		if res {
			t.Error("assert expected not to pass")
		}

		if ft.Buffer.String() != "regexp \"\\p\" didn't compile: error parsing regexp: invalid character class range: `\\p`" {
			t.Error("buffer expected to have error message, got:", ft.Buffer.String())
		}
	})
}

type FakeTesting struct {
	Buffer *bytes.Buffer
}

func (ft FakeTesting) Error(args ...interface{}) {
	ft.Buffer.WriteString(fmt.Sprint(args...))
}

func (ft FakeTesting) Errorf(format string, args ...interface{}) {
	ft.Buffer.WriteString(fmt.Sprintf(format, args...))
}

func (ft FakeTesting) Helper() {}
