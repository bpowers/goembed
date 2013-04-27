gorpi - Go library for Raspberry Pi embedded dev
================================================

The Raspberry Pi (RPi) is a great little linux board with readily
available GPIO, SPI, i2c and UART connectivity.  This is a library
designed to make controlling things using Go on a RPi easy and fun!

Reading a K-type thermocouple using a
[MAX31855](http://www.adafruit.com/products/269) looks like this:

```Go
package main

import (
	"github.com/bpowers/gorpi/device"
	"github.com/bpowers/gorpi/spi"
	"log"
)

func main() {
	tc1, err := device.Max31855(spi.Path(0, 0))
	if err != nil {
		log.Fatalf("devicesMax31855(spi.Path(0, 0)): %s\n", err)
	}
	defer tc1.Close()

	temp, err := tc1.Read()
	if err != nil {
		log.Fatalf("tc1.Read(): %s\n", err)
	}
	log.Printf("temp: %.2f°C (%.2f°F)\n", temp, temp*1.8+32)
}
```

Supported systems
-----------------

I am developing this on a Fedora 18 install which uses a
[rpi-3.6.y](https://github.com/raspberrypi/linux/commits/rpi-3.6.y)
kernel.  I would hope that it Just Works on other Linux distros with a
similar kernel, but if you run into issues I am happy to look into
them.

Similarly, I've tried to keep Linux-specific functionality in
*_linux.go files, I would hope that this code would be easy to port to
BSD, but I don't have plans to do that personally in the near future
unless there is a lot of demand.  Patches welcome ;)

license
-------

gorpi is offered under the MIT license, see LICENSE for details.
