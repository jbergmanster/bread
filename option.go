package bread

import (
	"math"
)

type OptType int

const (
	Call = iota
	Put
)

func Price(equityPrice, strikePrice float64, optType OptType) float64 {
	return math.Pow(equityPrice, 2)
}