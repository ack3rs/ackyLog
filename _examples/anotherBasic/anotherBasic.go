package main

import (
	myCustomLogger "github.com/acky666/ackyLog"
)

func main() {

	myCustomLogger.SHOWDEBUG = false
	myCustomLogger.DEBUG("This is never displayed %d %v", 1, "Bob")

}
