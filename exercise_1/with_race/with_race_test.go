package with_race

import "testing"

func TestIncrementWithRace(t *testing.T) {
	t.Run("counter must be incremented", func(t *testing.T) {
		got := IncrementWithRace()
		t.Logf("counter: %d", got)
	})
}
