package main

import (
	"github.com/megamsys/seru/cmd"
	"gopkg.in/check.v1"
)

func (s *S) TestNewSubdomainInfo(c *check.C) {
	desc := `creates a subdomain in the domain as it exists in the DNS service.`

	expected := &cmd.Info{
		Name:    "create",
		Usage:   `create -a <accesskey> -s <secretid> -d <domain name, default:megam.co> -u <subdomain> -i <ipaddress>`,
		Desc:    desc,
		MinArgs: 0,
	}
	command := NewSubdomain{}
	c.Assert(command.Info(), check.DeepEquals, expected)
}


func (s *S) TestDeleteSubdomainInfo(c *check.C) {
	desc := `deletes a subdomain.domain as it exists in the DNS service.`

	expected := &cmd.Info{
		Name:    "delete",
		Usage:   `delete -a <accesskey> -s <secretid> -d <domain name, default:megam.co> -s <subdomain>`,
		Desc:    desc,
		MinArgs: 0,
	}
	command := DeleteSubdomain{}
	c.Assert(command.Info(), check.DeepEquals, expected)
}
