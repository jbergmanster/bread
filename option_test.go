package bread

import "testing"

func TestPrice(t *testing.T) {
	const out = 5
	if x := Price(10, 10, Call); x != out {
		t.Errorf("Price of option is wrong.")
	}
}
