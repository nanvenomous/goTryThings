package sortstructs

import (
	"sort"
	"testing"
)

type NullableAgePerson struct {
	Age  *int
	Name string
}

var nlPeople []NullableAgePerson

func equalSlices(one []NullableAgePerson, two []NullableAgePerson) bool {
	for i := range one {
		if one[i].Name != two[i].Name {
			return false
		}
	}
	return true
}
func nlSlicesEqual(t *testing.T, one []NullableAgePerson, two []NullableAgePerson) {
	if (one == nil) || (two == nil) {
		t.Error("an array was nil")
	}
	if len(one) != len(two) {
		t.Error("array lengths not equal")
	}
	if !equalSlices(one, two) {
		t.Error("failed element-wise equality")
	}
}

func intPtr(num int) *int { return &num }

func setupNullable(t *testing.T) {
	nlPeople = []NullableAgePerson{
		{intPtr(40), "Diane"},
		{intPtr(15), "Katie"},
		{intPtr(41), "Margie"},
		{nil, "Emma"},
		{nil, "Colleen"},
		{intPtr(16), "Matt"},
	}
}

func TestSortByNullableAge(t *testing.T) {
	setupNullable(t)
	expected := []NullableAgePerson{
		{nil, "Colleen"},
		{nil, "Emma"},
		{intPtr(15), "Katie"},
		{intPtr(16), "Matt"},
		{intPtr(40), "Diane"},
		{intPtr(41), "Margie"},
	}
	sort.Slice(nlPeople[:], func(i, j int) bool {
		if (nlPeople[i].Age == nil) && (nlPeople[j].Age == nil) {
			return nlPeople[i].Name < nlPeople[j].Name
		} else if nlPeople[i].Age == nil || nlPeople[j].Age == nil {
			return nlPeople[i].Age == nil
		}
		return *nlPeople[i].Age < *nlPeople[j].Age
	})
	nlSlicesEqual(t, nlPeople, expected)
}
