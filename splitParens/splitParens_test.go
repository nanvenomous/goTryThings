package splitParens

import (
	"testing"
)

func TestSplitRuneArray(t *testing.T) {
	mock_ra := []rune("1234")
	ra1, ra2 := splitRuneArray(mock_ra, 2)
	if !(string(ra1) == "12") {
		t.Error("incorrect split")
	}
	if !(string(ra2) == "34") {
		t.Error("incorrect split")
	}
}

func TestNumberOfSplits1(t *testing.T) {
	mock_1 := "][(?"
	actual := numberOfSplits(mock_1)
	if !(actual == 1) {
		t.Error("returned incorrect number of splits", actual)
	}
}

func TestNumberOfSplits2(t *testing.T) {
	mock_1 := "(][)(?]["
	actual := numberOfSplits(mock_1)
	if !(actual == 2) {
		t.Error("returned incorrect number of splits", actual)
	}
}
