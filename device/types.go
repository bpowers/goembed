// Copyright 2013 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

type Thermocouple interface {
	Read() (Celsius, error)
	Precision() Celsius
	Close() error
}
