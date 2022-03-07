package itemdata

import (
	"errors"
	"fmt"
)

// ErrItemCodeNotFound is returned, when ItemFromString couldn't found appropriate code
var ErrItemCodeNotFound = errors.New("item code doesn't exist")

// ItemCodeFromStringWithError calls ItemCodeFromString but doesn't panic when string wasn't found.
func ItemCodeFromStringWithError(c string) (ItemCode, error) {
	var shouldReturnError bool
	// this should check if generated by string2enum function panics what means
	// that item code `c` doesn't exist
	func() {
		defer func() {
			if r := recover(); r != nil {
				shouldReturnError = true
			}
		}()

		ItemCodeFromString(c)
	}()

	if shouldReturnError {
		return ItemCodeNotFound, fmt.Errorf("getting wrapped item code for %s: %w", c, ErrItemCodeNotFound)
	}

	return ItemCodeFromString(c), nil
}
