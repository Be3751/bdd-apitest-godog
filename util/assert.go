package util

import (
	"fmt"
)

func AssertEqual(expected interface{}, actual interface{}) (err error) {
	if expected != actual {
		return fmt.Errorf("Not equal: the expected value was %v, but actually %v", expected, actual)
	}
	return nil
}
