goembed - Go library for embedded development (like the Raspberry Pi)
=====================================================================

The Raspberry Pi (RPi) is a great little linux board with readily
available GPIO, SPI, i2c and UART connectivity.  This is a library
designed to make controlling things using Go on a RPi easy and fun!

Reading a K-type thermocouple using a
[MAX31855](http://www.adafruit.com/products/269) looks like this:

```Go
package main

import (
	_ "github.com/bpowers/goembed/arch/raspberrypi"
	"github.com/bpowers/goembed/device"
	"github.com/bpowers/goembed/platform"
	"log"
)

const LEDPin = 22

func main() {
	maxSPI, err := platform.NewSPIPair(0, 0)
	if err != nil {
		log.Fatalf("platform.NewSPIPair(0, 0): %s\n", err)
	}

	tc1, err := device.Max31855(maxSPI)
	if err != nil {
		log.Fatalf("devices.NewMax31855(): %s", err)
	}
	defer tc1.Close()

	temp, err := tc1.Read()
	log.Printf("on startup, temp is %v (err: %v)", temp, err)

	pin, err := platform.OpenGPIO(LEDPin, platform.GPOutput)
	if err != nil {
		log.Fatalf("platform.OpenGPIO(LEDPin, platform.GPOutput): %s", err)
	}
	defer pin.Close()

	// turn on our LED
	pin.Write(1)

	// ...
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
