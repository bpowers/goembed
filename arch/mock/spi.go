// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package mock

import (
	"fmt"
	"github.com/bpowers/goembed/platform"
)

func spiUnimplemented(bus, slave int) (platform.SPIPair, error) {
	return nil, fmt.Errorf("mock SPI unimplemented")
}

func init() {
	platform.NewSPIPair = spiUnimplemented
}
