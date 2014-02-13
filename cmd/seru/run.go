package main

import (
	//"github.com/crowdmob/goamz/aws/route53"
	"github.com/indykish/seru/cmd"
	"launchpad.net/gnuflag"
)

type NewSubdomain struct {
	fs        *gnuflag.FlagSet
	accesskey string
	secretid  string
	domain    string
	subdomain string
}

func (g *NewSubdomain) Info() *cmd.Info {
	desc := `creates a subdomain in the domain as it exists in the DNS service.`
	return &cmd.Info{
		Name:    "create",
		Usage:    `create -A <aws_accesskey> -K <aws_secretid> -d <domain name, default:megam.co> 
		-s <subdomain> `,
		Desc:    desc,
		MinArgs: 4,
	}
}

func (c *NewSubdomain) Run(context *cmd.Context) error {
	//put code to call Route53 api.
	//https://github.com/karlentwistle/routemaster/blob/master/routemaster.go
	return nil
}

func (c *NewSubdomain) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("seru", gnuflag.ExitOnError)
		c.fs.StringVar(&c.accesskey, "accesskey", "", "AWS Accesskey")
		c.fs.StringVar(&c.accesskey, "A", "", "AWS Accesskey")
		c.fs.StringVar(&c.secretid, "secretid", "", "AWS Secretid")
		c.fs.StringVar(&c.secretid, "K", "", "AWS Secretid")
		c.fs.StringVar(&c.subdomain, "subdomain", "", "subdomain name")
		c.fs.StringVar(&c.subdomain, "s", "", "subdomain name")
		c.fs.StringVar(&c.domain, "domain", "megam.co", "domain name, this needs to preexist in the DNS service. Default : megam.co")
		c.fs.StringVar(&c.domain, "d", "megam.co", "domain name, this needs to preexist in the DNS service")
	}
	return c.fs
}

type DeleteSubdomain struct {
	fs        *gnuflag.FlagSet
	accesskey string
	secretid  string
	domain    string
	subdomain string
}

func (g *DeleteSubdomain) Info() *cmd.Info {
	desc := `deletes a subdomain.domain as it exists in the DNS service.`
	return &cmd.Info{
		Name:    "delete",
		Usage:    `delete -A <aws_accesskey> -K <aws_secretid> -d <domain name, default:megam.co> 
		-s <subdomain> `,
		Desc:    desc,
		MinArgs: 4,
	}
}

func (c *DeleteSubdomain) Run(context *cmd.Context) error {
	return nil
}

func (c *DeleteSubdomain) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("seru", gnuflag.ExitOnError)
		c.fs.StringVar(&c.accesskey, "accesskey", "", "AWS Accesskey")
		c.fs.StringVar(&c.accesskey, "A", "", "AWS Accesskey")
		c.fs.StringVar(&c.secretid, "secretid", "", "AWS Secretid")
		c.fs.StringVar(&c.secretid, "K", "", "AWS Secretid")
		c.fs.StringVar(&c.subdomain, "subdomain", "", "subdomain name")
		c.fs.StringVar(&c.subdomain, "s", "", "subdomain name")
		c.fs.StringVar(&c.domain, "domain", "megam.co", "domain name, this needs to preexist in the DNS service. Default : megam.co")
		c.fs.StringVar(&c.domain, "d", "megam.co", "domain name, this needs to preexist in the DNS service")
	}
	return c.fs
}
