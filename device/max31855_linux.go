// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/bpowers/gorpi/spi"
	"os"
)

type max31855Error string

func (m max31855Error) Error() string {
	return string(m)
}

const (
	Max31855ShortCircuitVCC = max31855Error("Max31855: short-circuited to VCC")
	Max31855ShortCircuitGround = max31855Error("Max31855: short-circuited to VCC")
	Max31855OpenConn = max31855Error("Max31855: open (no connections)")
)

type max31855 struct {
	f *os.File
}

type max31855Reply struct {
	TCData int16
	RJData int16
	Fault  bool // any of SCV, SCG or OC are non-zero
	SCV    bool // short-circuited to VCC (fault)
	SCG    bool // short-circuited to GND (fault)
	OC     bool // open (no connections) (fault)
}

func Max31855(path spi.SPIPath) (Thermocouple, error) {
	f, err := os.OpenFile(string(path), os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("os.OpenFile('%s', os.O_RDWR, 0)", path)
	}
	return &max31855{f}, nil
}

func (m *max31855) read() (max31855Reply, error) {
	// Reads are 4-bytes
	buf := make([]byte, 4)
	var reply max31855Reply

	if err := spi.Transaction(m.f, nil, buf); err != nil {
		return reply, fmt.Errorf("spi.Transaction(%v, nil, buf): %s", m.f, err)
	}

	reply.TCData = int16(buf[0]<<6|buf[1]>>2)
	reply.RJData = int16(buf[2]<<4|buf[3]>>4)
	reply.Fault = (buf[1]>>7) == 1
	reply.SCV = ((buf[3]&4)>>2) == 1
	reply.SCG = ((buf[3]&2)>>1) == 1
	reply.OC = (buf[3]&1) == 1

	return reply, nil
}

func (m *max31855) Read() (Celsius, error) {
	reply, err := m.read()
	if err != nil {
		return -1, err
	}

	switch {
	case reply.SCV:
		return -1, Max31855ShortCircuitVCC
	case reply.SCG:
		return -1, Max31855ShortCircuitGround
	case reply.OC:
		return -1, Max31855OpenConn
	}

	// the high 14 bits of the result contain 4x the Celsius temp
	return Celsius(m.Precision() * Celsius(reply.TCData)), nil
}

func (m *max31855) Precision() Celsius {
	return Celsius(.25)
}

func (m *max31855) Close() error {
	return m.f.Close()
}
