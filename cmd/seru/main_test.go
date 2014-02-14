package main

import (
	"github.com/indykish/seru/cmd"
	"launchpad.net/gocheck"
)

func (s *S) TestCommandsFromBaseManagerAreRegistered(c *gocheck.C) {
	baseManager := cmd.BuildBaseManager("megam", version, header)
	manager := buildManager("megam")
	for name, instance := range baseManager.Commands {
		command, ok := manager.Commands[name]
		c.Assert(ok, gocheck.Equals, true)
		c.Assert(command, gocheck.FitsTypeOf, instance)
	}
}

/*func (s *S) TestNewSubdomainIsRegistered(c *gocheck.C) {
	manager := buildManager("megam")
	create, ok := manager.Commands["create"]
	c.Assert(ok, gocheck.Equals, true)
	c.Assert(create, gocheck.FitsTypeOf, &NewSubdomain{})
}

func (s *S) TestDeleteSubdomainIsRegistered(c *gocheck.C) {
	manager := buildManager("megam")
	remove, ok := manager.Commands["delete"]
	c.Assert(ok, gocheck.Equals, true)
	c.Assert(remove, gocheck.FitsTypeOf, &DeleteSubdomain{})
}*/