package bread

import (
	"math"
)

// Position is whether this is a call or put.
type Position int

const (
	Call Position = iota
	Put
)

// Exercise determines when the option can be exercised.
type Exercise int

const (
	American Exercise = iota
	European
)

// Instrument is the type of option being priced.
type Instrument struct {
	P  Position
	E  Exercise
	Eq float64 // equity price at purchase
	S  float64 // strike price
}

// BinSimParams are parameters use to price an option using the Binomial pricing model.
type BinSimParams struct {
	R float64 // interest rate
	U float64 // up step
	D float64 // down step
	N int     // number of time steps
}

// BinomialPrice prices option using a Binomial tree.
func BinomialPrice(inst Instrument, params BinSimParams) float64 {
	a := make([]float64, params.N)
	for i, _ := range a {
		t := N - i
		h := i
		a[i] = math.Pow(params.U, h) * math.pow(params.D, t)
	}
	return math.Pow(equityPrice, 2)
}
