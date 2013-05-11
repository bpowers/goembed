// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package mock

import (
	"fmt"
	"github.com/bpowers/goembed/platform"
	"log"
)

type GPIO struct {
	state int8
	pin int
	dir platform.GPIODir
}

func NewGPIO(pin int, dir platform.GPIODir) (platform.GPIO, error) {
	return &GPIO{-1, pin, dir}, nil
}

func (g *GPIO) Read() (byte, error) {
	if g.dir & platform.GPInput != platform.GPInput {
		return 0, fmt.Errorf("read of output pin %d", g.pin)
	}
	if g.state == -1 {
		log.Printf("WARNING: uninitialized read of GPIO %d", g.pin)
	}

	if g.state == 1 {
		return 1, nil
	} else {
		return 0, nil
	} 
}

func (g *GPIO) Write(b byte) error {
	if g.dir & platform.GPOutput != platform.GPOutput {
		return fmt.Errorf("write to input pin %d", g.pin)
	}
	g.state = int8(b & 0x01)
	return nil
}

func (g *GPIO) Dir() platform.GPIODir {
	return g.dir
}

func (g *GPIO) Close() error {
	return nil
}

func init() {
	platform.OpenGPIO = NewGPIO
}
