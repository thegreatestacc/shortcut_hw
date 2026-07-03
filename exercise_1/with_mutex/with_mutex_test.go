package with_mutex

import "testing"

func TestIncrementWithMutex(t *testing.T) {
	t.Run("counter must be 100000", func(t *testing.T) {
		want := 100000
		got := IncrementWithMutex()

		if got != want {
			t.Errorf("increment value \n got %d, want %d", got, want)
		}
	})
}
