// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package raspberrypi

import (
	"fmt"
	"github.com/bpowers/goembed/platform"
	"os"
	"strconv"
)

type rpiGPIO struct {
	f *os.File
	pin int
	dir platform.GPIODir
}

func pinPath(pin int, file string) string {
	return fmt.Sprintf("/sys/class/gpio/gpio%d/%s", pin, file)
}

func exportPin(pin int) error {
	// if the pin already exists, our job is already done.
	if _, err := os.Stat(pinPath(pin, "")); err == nil {
		return nil
	}

	f, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("os.Create(/sys/class/gpio/export): %s", err)
	}
	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(pin))
	return err
}

func openRPiGPIO(pin int, dir platform.GPIODir) (platform.GPIO, error) {
	var err error
	if err = exportPin(pin); err != nil {
		return nil, fmt.Errorf("exportPin(%d): %s", pin, err)
	}

	dirF, err := os.OpenFile(pinPath(pin, "direction"), os.O_WRONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("os.Create(%s): %s",
			pinPath(pin, "direction"), err)
	}
	defer dirF.Close()

	if _, err = dirF.WriteString(dir.String());  err != nil {
		return nil, fmt.Errorf("dirF.WriteString(%s): %s",dir.String(), err)
	}

	var mask int
	if dir == platform.GPInput || dir == platform.GPBidi {
		mask |= os.O_RDONLY
	}
	if dir == platform.GPOutput || dir == platform.GPBidi {
		mask |= os.O_WRONLY
	}

	valF, err := os.OpenFile(pinPath(pin, "value"), mask, 0)
	if err != nil {
		return nil, fmt.Errorf("os.Create(%s): %s", pinPath(pin, "value"), err)
	}

	return &rpiGPIO{valF, pin, dir}, nil
}

func (r *rpiGPIO) Read() (byte, error) {
	buf := []byte{0}
	_, err := r.f.Read(buf)
	// FIXME: make this more robust
	return (buf[0]-'0')&0x01, err
}

func (r *rpiGPIO) Write(b byte) error {
	buf := []byte{0}
	buf[0] = '0' + (b&0x01)
	_, err := r.f.Write(buf)
	return err
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
