package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	if len(e.Errors) <= 0 {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(len(e.Errors)))
	sb.WriteString(" errors occured:\n")

	for _, err := range e.Errors {
		sb.WriteString("\t* ")
		sb.WriteString(err.Error())
	}

	sb.WriteString("\n")

	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	if me, ok := err.(*MultiError); ok {
		me.Errors = append(me.Errors, errs...)
		return me
	}

	me := &MultiError{
		Errors: make([]error, 0),
	}

	if err != nil {
		me.Errors = append(me.Errors, err)
	}

	me.Errors = append(me.Errors, errs...)

	return me
}

func main() {
	var err error
	err = errors.New("error 0")
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))
	fmt.Println(err.Error())
}
