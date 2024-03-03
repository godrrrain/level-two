package main

import (
	"reflect"
	"testing"
)

func TestFindAnagram(t *testing.T) {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стакан"}
	expect := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	res := findAnagram(&input)

	if !reflect.DeepEqual(*res, expect) {
		t.Errorf("expected %v \nbut got %v", expect, *res)
	}
}
