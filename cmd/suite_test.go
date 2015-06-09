package cmd

import (
	"bytes"
	"gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	stdin   *os.File
	recover []string
}

var _ = check.Suite(&S{})
var manager *Manager


func (s *S) SetUpTest(c *check.C) {
	var stdout, stderr bytes.Buffer
	manager = NewManager("seru", "0.1", "", &stdout, &stderr, os.Stdin)
	var exiter recordingExiter
	manager.e = &exiter
}
