package main

import (
	"bytes"
	"github.com/indykish/seru/cmd"
	"launchpad.net/gocheck"
	"os"
	"testing"
)

type S struct {
	recover []string
}

var _ = gocheck.Suite(&S{})
var manager *cmd.Manager

func Test(t *testing.T) { gocheck.TestingT(t) }

func (s *S) SetUpTest(c *gocheck.C) {
	var stdout, stderr bytes.Buffer
	manager = cmd.NewManager("seru", version, header, &stdout, &stderr, os.Stdin)
}
