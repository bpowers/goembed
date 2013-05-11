// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package platform provides standard interfaces and creation
// functions for common embedded buses and IO methods.
package platform

// SPIPair represents the connection between an SPI bus master and a
// specific slave.  Different slaves connected to the same SPI master
// have distinct SPIPairs.
type SPIPair interface {
	// Transaction sends the data in the write buffer to the
	// slave, and fills the read buffer with len(read) bytes of
	// data.  Either buffer is allowed to be nil, but not both.
	Transaction(write, read []byte) error
	Close() error
}

// NewSPIPair is the platform-specific SPIPair creation function, it
// is installed by importing a specific arch in main, like:
//
//     import _ "github.com/bpowers/goembed/arch/raspberry-pi"
var NewSPIPair func(bus, slave int) (SPIPair, error)

type GPIODir int // the direction (input or output) of the IO pin

func (d GPIODir) String() string {
	switch d {
	case GPInput:
		return "in"
	case GPOutput:
		return "out"
	case GPBidi:
		return "bidi"
	}
	panic("unknown GPIODir")
}

const (
	GPInput GPIODir = 1 << 0
	GPOutput GPIODir = 1 << 1
	GPBidi GPIODir = GPInput | GPOutput
)

// GPIO represents an individual pin opened for either input or output
type GPIO interface {
	Read() (byte, error) // only the lowest bit (0x01) will possibly be set
	Write(b byte) error // the input, b, will be bitwise-and'ed with 0x01
	Dir() GPIODir
	Close() error
}

// OpenGPIO is the platform-specific way to gain access to a GPIO pin.
var OpenGPIO func(pin int, dir GPIODir) (GPIO, error)
