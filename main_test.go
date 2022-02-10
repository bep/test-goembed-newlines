package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestEmbed(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	c.Assert(embedHello, qt.Equals, fileHello())
	// c.Assert(embedHello, qt.Equals, "\nHello.\n")
}
