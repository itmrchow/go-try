package main

import (
	"fmt"

	"github.com/pkg/errors"
	// "github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := fn()

	// causeErr := errors.Cause(err)

	println(err.Error())
	println(errors.Unwrap(err).Error())
	// println(causeErr.Error())
}

func fn() error {
	e1 := errors.New("1")
	e2 := fmt.Errorf("2"+":%w", e1)
	// e2 := errors.Wrap(e1, "inner")
	// e3 := errors.Wrap(e2, "middle")
	e4 := fmt.Errorf("3"+":%w", e2)
	return e4
}
