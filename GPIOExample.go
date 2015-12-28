package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/davecheney/gpio"
	//"github.com/davecheney/gpio/rpi"
)
func main() {
	// Create an Array
	pins := [...]int{2, 3, 4, 17, 27, 22, 10, 9, 11, 7, 8, 25, 24, 23, 18}
	// Loop through each element in that array
	for _,z:=range pins {
		// call function setPin with the value of z
		setPin(z)
	}
}
func setPin(p int) {
	// set GPIO to output mode
	pin, err := gpio.OpenPin(p, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	// turn the led off on exit
	// this is if you press CTRL-C to exit the program early
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nClearing and unexporting the pin.\n")
			pin.Clear()
			pin.Close()
			os.Exit(0)
		}
	}()

		fmt.Printf("Setting GPIO %v Out\n", p)
		pin.Set()
		time.Sleep(3 * time.Second)
		fmt.Printf("Setting GPIO %v In\n", p)
		pin.Clear()
		pin.Close()
		time.Sleep(3 * time.Second)
}
