package main

import (
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestEmbed(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	c.Assert(embedHello, qt.Equals, fileHello())

	c.Assert(strings.ReplaceAll(embedHello, "\r\n", "\n"), qt.Equals, "\nHello.\n")
}
