package ackyLog

import "testing"

//TODO:
func TestLog(t *testing.T) {
	// you can turn them off completely.
	SHOWCOLOURS = false

	WARNING("This is a WARNING Message %d", 1)

}
