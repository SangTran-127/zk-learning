package main

import (
	"learning-zk/circuit"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func main() {
	var c circuit.SquareCircuit

	constraintSystem, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &c)

	if err != nil {
		panic(err)
	}

	assignment := circuit.SquareCircuit{
		X: 3,
		Y: 9,
	}

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	if err != nil {
		panic(err)
	}

	// Verifier need the public witness
	// Private witness for prover
	publicWitness, _ := witness.Public()

	provingKey, verificationKey, err := groth16.Setup(constraintSystem)

	if err != nil {
		panic(err)
	}

	// Prover will need the whole witness
	proof, err := groth16.Prove(constraintSystem, provingKey, witness)

	if err != nil {
		panic(err)
	}
	// Verifier just need the public witness
	verifyError := groth16.Verify(proof, verificationKey, publicWitness)

	if verifyError != nil {
		panic(verifyError)
	} else {
		println("Verified success")
	}

}
