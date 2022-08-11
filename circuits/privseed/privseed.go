// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package privseed

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

// Circuit defines a simple circuit
// x**3 + x + 5 == y
type Circuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	KeyID      frontend.Variable
	KeyAddress frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints
// gen_keypair(hint(password)).pub  == pub
func (circuit *Circuit) Define(api frontend.API) error {
	res, err := api.Compiler().NewHint(GetSeedFromID, 1, circuit.KeyID)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	seedInt := res[0]

	res, err = api.Compiler().NewHint(GenerateKey, 1, seedInt)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	pubkey := res[0]

	res, err = api.Compiler().NewHint(GetAddress, 1, pubkey)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	address := res[0]
	api.AssertIsEqual(address, circuit.KeyAddress)
	return nil
}

func GetAddress(curveID ecc.ID, inputs, outputs []*big.Int) error {
	return nil
}

func GenerateKey(curveID ecc.ID, inputs, outputs []*big.Int) error {
	return nil
}

func GetSeedFromID(curveID ecc.ID, inputs, outputs []*big.Int) error {
	return nil
}
