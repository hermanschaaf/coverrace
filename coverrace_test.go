package coverrace

import "testing"

func TestCoverRace(t *testing.T) {
	got := add100()
	if got != 100 {
		t.Errorf("got %d, want %d", got, 100)
	}
}
