package sortstructs

import (
	"fmt"
	"sort"
	"testing"
)

type Person struct {
	Age  int
	Name string
}

var people []Person

func slicesEqual(t *testing.T, one []Person, two []Person) {
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

func setup(t *testing.T) {
	people = []Person{{12, "Emma"}, {40, "Diane"}, {15, "Katie"}}
}

func TestSortByAge(t *testing.T) {
	setup(t)
	fmt.Println("people: ", people)
	expected := []Person{{12, "Emma"}, {15, "Katie"}, {40, "Diane"}}
	sort.Slice(people[:], func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	slicesEqual(t, people, expected)
}

func TestSortByName(t *testing.T) {
	setup(t)
	expected := []Person{{40, "Diane"}, {12, "Emma"}, {15, "Katie"}}
	sort.Slice(people[:], func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	slicesEqual(t, people, expected)
}
