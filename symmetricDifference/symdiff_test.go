package symdiff

import (
	"testing"
)

func slicesEqual(t *testing.T, one []string, two []string) {
	if (one == nil) || (two == nil) {
		t.Error("an array was nil")
	}
	if len(one) != len(two) {
		t.Error("array lengths not equal")
	}
	for i := range one {
		if one[i] != two[i] {
			t.Errorf("inequality at %d", i)
		}
	}
}

func TestPrependArgument(t *testing.T) {
	sliceExtra := []string{"0ne", "two", "three"}
	slice := []string{"0ne"}
	extras := []string{"two", "three"}
	slicesEqual(t, extras, SymmetricDifference(sliceExtra, slice))
}
