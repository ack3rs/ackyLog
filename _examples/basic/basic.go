package main

import (
	"errors"
	"time"

	myCustomLogger "github.com/acky666/ackyLog"
)

func main() {

	// Disable ALL colours. Normally wrap this around some environmental varibles.
	// for example:  the AWS LightSail container service or if your sending the logs
	// to Elastic Search, the colours in the logs are an issue for readability so
	// you can turn them off completely.
	myCustomLogger.SHOWCOLOURS = false

	myCustomLogger.WARNING("This is a WARNING Message %d", 1)

	// Uses the awesome Spew library
	myCustomLogger.SPEW("Bob")

	// I know it's the not the done, thing, but i normally use this libray like l.INFO
	// and l.DEBUG etc, and I find the functions easier to read if they are upper case

	myCustomLogger.SHOWCOLOURS = true // This is the Default
	myCustomLogger.WARNING("This is a WARNING Message %d", 1)
	myCustomLogger.DEBUG("This is a DEBUG Message %d %v", 1, "Bob")
	myCustomLogger.INFO("This is a INFO Message %d", 234)
	myCustomLogger.ERROR("This is a ERROR %s", errors.New("My Big Error"))
	myCustomLogger.RAW("This is a RAW Message with NO formatting")

	t := func() {
		// Do something on a GoRoutine
		timed := myCustomLogger.TIMED("Timed in a GoRoutine")
		defer timed()
		time.Sleep(2 * time.Second)
		myCustomLogger.DEBUG("This is a DEBUG Message %d %v", 1, "Bob")
	}
	go t()

	timed := myCustomLogger.TIMED("This is Timed")
	time.Sleep(5 * time.Second)
	timed()

}
