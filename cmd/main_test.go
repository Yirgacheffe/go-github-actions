package main

import (
	"bytes"
	"strings"

	"testing"

	"github.com/matryer/is"
)

func Test_WithNames(t *testing.T) {
	is := is.New(t)

	args := []string{"greeter", "David", "Kat", "Jon", "Natalie", "Mark"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "Hi David"))
	is.True(strings.Contains(out, "Hi Kat"))
	is.True(strings.Contains(out, "Hi Jon"))
	is.True(strings.Contains(out, "Hi Natalie"))
	is.True(strings.Contains(out, "Hi Mark"))
}

func Test_NoNames(t *testing.T) {
	is := is.New(t)

	args := []string{"greeter", "anabel"}
	var stdout bytes.Buffer

	var err error
	err = run(args, &stdout)
	is.NoErr(err)

	err = runWithFlags([]string{"greeter", "-v", "-f=json"}, &stdout)
	is.NoErr(err)
}
