// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/bpowers/goembed/platform"
)

type TRIACDimmer struct {
	ZeroCross platform.GPIO
	Output    platform.GPIO
	closeChan chan bool
}

func NewTRIACDimmer(zeroCross, out platform.GPIO) (*TRIACDimmer, error) {
	if zeroCross.Dir() & platform.GPInput != platform.GPInput {
		return nil, fmt.Errorf("zero-cross GPIO not setup for input")
	}
	if out.Dir() & platform.GPOutput != platform.GPOutput {
		return nil, fmt.Errorf("output GPIO not setup for output")
	}

	dim := &TRIACDimmer{
		ZeroCross: zeroCross,
		Output: out,
	}
	return dim, nil
}

func (t *TRIACDimmer) Close() error {
	return nil
}

func (t *TRIACDimmer) worker() {
outer:
	for {
		select {
		case <- t.closeChan:
			break outer
		}
	}
	fmt.Printf("worker closing")
}
