package main

import (
	"github.com/megamsys/seru/cmd"
	"gopkg.in/check.v1"
)

func (s *S) TestCommandsFromBaseManagerAreRegistered(c *check.C) {
	baseManager := cmd.BuildBaseManager("megam", version, header)
	manager := buildManager("megam")
	for name, instance := range baseManager.Commands {
		command, ok := manager.Commands[name]
		c.Assert(ok, check.Equals, true)
		c.Assert(command, check.FitsTypeOf, instance)
	}
}

/*func (s *S) TestNewSubdomainIsRegistered(c *check.C) {
	manager := buildManager("megam")
	create, ok := manager.Commands["create"]
	c.Assert(ok, check.Equals, true)
	c.Assert(create, check.FitsTypeOf, &NewSubdomain{})
}

func (s *S) TestDeleteSubdomainIsRegistered(c *check.C) {
	manager := buildManager("megam")
	remove, ok := manager.Commands["delete"]
	c.Assert(ok, check.Equals, true)
	c.Assert(remove, check.FitsTypeOf, &DeleteSubdomain{})
}*/
