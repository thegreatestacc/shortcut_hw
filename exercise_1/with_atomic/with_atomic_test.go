package with_atomic

import "testing"

func TestIncrementWithAtomic(t *testing.T) {
	t.Run("counter must be 100000", func(t *testing.T) {
		want := int64(100000)
		got := IncrementWithAtomic()

		if got != want {
			t.Errorf("increment value \n got %d, want %d", got, want)
		}
	})
}
