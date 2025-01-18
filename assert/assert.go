package assert

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/mishankov/testman/internal/interfaces"
)

func True(t interfaces.TB, condition bool) bool {
	t.Helper()

	if !condition {
		t.Error("condition expected to be true")
		return false
	}

	return true
}

func Equal[T comparable](t interfaces.TB, got, want T) bool {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
		return false
	}

	return true
}

func DeepEqual[T any](t interfaces.TB, got, want T) bool {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
		return false
	}

	return true
}

func Contains(t interfaces.TB, str, substr string) bool {
	t.Helper()

	if !strings.Contains(str, substr) {
		t.Errorf("expected %q to contain %q", str, substr)
		return false
	}

	return true
}

func Regex(t interfaces.TB, got, wantRegex string) bool {
	t.Helper()

	r, err := regexp.Compile(wantRegex)
	if err != nil {
		t.Errorf(`regexp "%v" didn't compile: %v`, wantRegex, err)
		return false
	}

	if !r.MatchString(got) {
		t.Errorf(`%q didn't match regexp "%v"`, got, wantRegex)
		return false
	}

	return true
}
