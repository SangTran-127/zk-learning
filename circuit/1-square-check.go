package circuit

import (
	"github.com/consensys/gnark/frontend"
)

// This circuit will prove x * x = y without telling what is x

type SquareCircuit struct {
	// Private by default
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *SquareCircuit) Define(api frontend.API) error {
	api.AssertIsEqual(api.Mul(circuit.X, circuit.X), circuit.Y)
	return nil
}
