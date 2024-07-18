package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("2 + 2", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectMsg(t, got, want)
	})

	t.Run("1 + 2", func(t *testing.T) {
		got := Add(1, 2)
		want := 1 + 2
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}
