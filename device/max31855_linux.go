// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/bpowers/gorpi/spi"
	"os"
)

type max31855 struct {
	f *os.File
}

func Max31855(path spi.SPIPath) (Thermocouple, error) {
	f, err := os.OpenFile(string(path), os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("os.OpenFile('%s', os.O_RDWR, 0)", path)
	}
	return &max31855{f}, nil
}

func (m *max31855) Read() (Celsius, error) {
	// Reads are 4-bytes
	buf := make([]byte, 4)

	if err := spi.Transaction(m.f, nil, buf); err != nil {
		return 0, fmt.Errorf("spi.Transaction(%v, nil, buf): %s", m.f, err)
	}

	// the high 14 bits of the result contain 4x the Celsius temp
	return Celsius(m.Precision() * Celsius(buf[0]<<6|buf[1]>>2)), nil
}

func (m *max31855) Precision() Celsius {
	return Celsius(.25)
}

func (m *max31855) Close() error {
	return m.f.Close()
}
