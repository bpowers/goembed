// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package device

import (
	"io"
)

type Celsius float64

type Thermocouple interface {
	Read() (Celsius, error)
	Precision() Celsius
	io.Closer
}
