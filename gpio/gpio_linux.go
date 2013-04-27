// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package devices

import (
	"fmt"
	"os"
	"sync"
)

var (
	pinsLock sync.RWMutex
	pins     = map[Pin]*GPIO{}
)

// direction (input or output) of the IO pin
type Direction int

// integer number of the pin
type Pin int

const (
	DirIn  Direction = 1
	DirOut Direction = 2
)

type GPIO struct {
	f   *os.File
	Dir Direction
	Pin Pin
}

// Returns an opened GPIO pin for input
func OpenInput(n Pin, initial int) (*GPIO, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Returns an opened GPIO pin for output
func OpenOutput(n Pin, initial int) (*GPIO, error) {
	return nil, fmt.Errorf("unimplemented")
}
