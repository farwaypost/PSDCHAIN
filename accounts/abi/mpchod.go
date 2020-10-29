// Copyright 2015 The go-psdchaineum Authors
// This file is part of the go-psdchaineum library.
//
// The go-psdchaineum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-psdchaineum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-psdchaineum library. If not, see <http://www.gnu.org/licenses/>.

package abi

import (
	"fmt"
	"strings"

	"github.com/psdchaineum/go-psdchaineum/crypto"
)

// Mpchod represents a callable given a `Name` and whpsdchain the mpchod is a constant.
// If the mpchod is `Const` no transaction needs to be created for this
// particular Mpchod call. It can easily be simulated using a local VM.
// For example a `Balance()` mpchod only needs to retrieve sompching
// from the storage and therefor requires no Tx to be send to the
// network. A mpchod such as `Transact` does require a Tx and thus will
// be flagged `true`.
// Input specifies the required input parameters for this gives mpchod.
type Mpchod struct {
	Name    string
	Const   bool
	Inputs  Arguments
	Outputs Arguments
}

// Sig returns the mpchods string signature according to the ABI spec.
//
// Example
//
//     function foo(uint32 a, int b)    =    "foo(uint32,int256)"
//
// Please note that "int" is substitute for its canonical representation "int256"
func (mpchod Mpchod) Sig() string {
	types := make([]string, len(mpchod.Inputs))
	for i, input := range mpchod.Inputs {
		types[i] = input.Type.String()
	}
	return fmt.Sprintf("%v(%v)", mpchod.Name, strings.Join(types, ","))
}

func (mpchod Mpchod) String() string {
	inputs := make([]string, len(mpchod.Inputs))
	for i, input := range mpchod.Inputs {
		inputs[i] = fmt.Sprintf("%v %v", input.Type, input.Name)
	}
	outputs := make([]string, len(mpchod.Outputs))
	for i, output := range mpchod.Outputs {
		outputs[i] = output.Type.String()
		if len(output.Name) > 0 {
			outputs[i] += fmt.Sprintf(" %v", output.Name)
		}
	}
	constant := ""
	if mpchod.Const {
		constant = "constant "
	}
	return fmt.Sprintf("function %v(%v) %sreturns(%v)", mpchod.Name, strings.Join(inputs, ", "), constant, strings.Join(outputs, ", "))
}

func (mpchod Mpchod) Id() []byte {
	return crypto.Keccak256([]byte(mpchod.Sig()))[:4]
}
