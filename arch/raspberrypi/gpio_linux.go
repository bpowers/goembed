// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package raspberrypi

import (
	"fmt"
	"github.com/bpowers/goembed/platform"
	"os"
)

type rpiGPIO struct {
	dir platform.GPIODir
	f *os.File
}

func openRPiGPIO(pin int, dir platform.GPIODir) (platform.GPIO, error) {
	return nil, fmt.Errorf("ni")
}

func (r *rpiGPIO) Read() (byte, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *rpiGPIO) Write(b byte) error {
	return fmt.Errorf("not implemented")
}

func (r *rpiGPIO) Dir() platform.GPIODir {
	return r.dir
}

func (r *rpiGPIO) Close() error {
	return r.f.Close()
}

func init() {
	platform.OpenGPIO = openRPiGPIO
}
