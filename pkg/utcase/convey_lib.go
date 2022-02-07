package utcase

import (
	"errors"
)

func example() error {
	return nil
}

var example2 = func() int {
	return 10
}

type c struct {
	Branch bool
}

func (c *c) example3() error {
	if c.Branch {
		return errors.New("bad branch")
	}
	return nil
}
