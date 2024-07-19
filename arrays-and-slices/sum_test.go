package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 1...5", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15
		assertCorrectMsg(t, got, want)
	})

	t.Run("sum of 1...10", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		want := 55
		assertCorrectMsg(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of several slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 9})
		want := []int{15, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sum of all excepts head", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{1, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
