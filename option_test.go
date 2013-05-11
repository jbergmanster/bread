package bread

import "testing"

func TestPrice(t *testing.T) {
	sim := BinSimParams{
		R: .25,
		U: 2.0,
		D: .5,
		N: 4}
	inst := Instrument{P: Call, E: European, Pr: 4, S: 6}
	out := 4.0
	if x := sim.Price(inst); x != out {
		t.Errorf("Price of option is wrong. Want %f, Got %f", out, x)
	}
}
