package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	want := []int{0, 1}

	got := twoSum(input, target)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Excepted %q, got %q", want, got)
	}
}

func TestTwoSum2(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	want := []int{0, 1}

	got := twoSum2(input, target)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Excepted %q, got %q", want, got)
	}
}
