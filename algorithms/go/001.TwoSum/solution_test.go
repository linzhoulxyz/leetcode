package TwoSum

import (
	"testing"
	"reflect"
)

func TestTwoSum(t *testing.T) {
	input := []int{2,7,11,15}
	target := 9
	want := []int{0, 1}

	got := twoSum(input, target)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Excepted %q, got %q", want, got)
	}
}