package main

import (
	"github.com/indykish/seru/cmd"
	"launchpad.net/gocheck"
)

func (s *S) TestNewSubdomainInfo(c *gocheck.C) {
	desc := `creates a subdomain in the domain as it exists in the DNS service.`

	expected := &cmd.Info{
		Name:    "create",	
		Usage:    `seru create -A <aws_accesskey> -K <aws_secretid> -d <domain name, default:megam.co> 
		-s <subdomain> `,	
		Desc:    desc,
		MinArgs: 4,
	}
	command := NewSubdomain{}
	c.Assert(command.Info(), gocheck.DeepEquals, expected)
}


func (s *S) TestDeleteSubdomainInfo(c *gocheck.C) {
	desc := `deletes a subdomain.domain as it exists in the DNS service.`
	
	expected := &cmd.Info{
		Name:    "delete",
		Usage:    `seru delete -A <aws_accesskey> -K <aws_secretid> -d <domain name, default:megam.co> 
		-s <subdomain> `,		
		Desc:    desc,
		MinArgs: 4,
	}
	command := DeleteSubdomain{}
	c.Assert(command.Info(), gocheck.DeepEquals, expected)
}

