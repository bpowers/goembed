// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package raspberrypi

import (
	"fmt"
	"github.com/bpowers/goembed/platform"
	"os"
	"unsafe"
)

const (
	_SPI_IOC_MAGIC = 'k'
)

type rpiSPIPair struct {
	*os.File
}

func newSPIPair(bus, slave int) (platform.SPIPair, error) {
	devPath := fmt.Sprintf("/dev/spidev%d.%d", bus, slave)
	f, err := os.Create(devPath)
	if err != nil {
		return nil, fmt.Errorf("os.OpenFile('%s', os.O_RDWR, 0)", devPath)
	}
	return &rpiSPIPair{f}, nil
}

type _SPIIOTransaction struct {
	TXBuf       uint64
	RXBuf       uint64
	Len         uint32
	SpeedHz     uint32
	DelayUsecs  uint16
	BitsPerWord uint8
	CSChange    uint8
	Pad         uint32
}

const sizeof_SPIIOTransaction = 32 // bytes

func _SPI_IOC_MESSAGE(count int) int32 {
	return _IOW(_SPI_IOC_MAGIC, 0, count*sizeof_SPIIOTransaction)
}

func (s *rpiSPIPair) Transaction(write, read []byte) error {
	if write != nil && read != nil && len(write) != len(read) {
		return fmt.Errorf("write and read size mismatch (%d vs %d)",
			len(write), len(read))
	}
	var length uint32
	var wp, rp unsafe.Pointer
	if write != nil {
		wp = unsafe.Pointer(&write[0])
		length = uint32(len(write))
	}
	if read != nil {
		rp = unsafe.Pointer(&read[0])
		length = uint32(len(read))
	}
	trx := _SPIIOTransaction{
		TXBuf: uint64(uintptr(wp)),
		RXBuf: uint64(uintptr(rp)),
		Len:   length,
	}
	result := ioctl(s.Fd(), _SPI_IOC_MESSAGE(1), unsafe.Pointer(&trx))
	if result != 0 {
		return fmt.Errorf("ioctl result of %d", result)
	}
	return nil
}
