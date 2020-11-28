package sum

import "testing"

func TestSum(t *testing.T) {
	number := [...]int{1, 2, 3, 4, 5}
	got := Sum(number)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, number)
	}
}
