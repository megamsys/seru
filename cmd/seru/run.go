package main

import (
	"github.com/karlentwistle/route53"
	"github.com/indykish/seru/cmd"
	"launchpad.net/gnuflag"
	"log"
	"fmt"
	"os"
)

type NewSubdomain struct {
	fs        *gnuflag.FlagSet
	accesskey string
	secretid  string
	domain    string
	subdomain string
	ip        string
}

func (c *NewSubdomain) Info() *cmd.Info {
	desc := `creates a subdomain in the domain as it exists in the DNS service.`
	return &cmd.Info{
		Name:    "create",
		Usage:   `create -a <accesskey> -s <secretid> -d <domain name, default:megam.co> -u <subdomain> -i <ipaddress>`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (c *NewSubdomain) Run(context *cmd.Context) error {
	aws := auth(c.accesskey, c.secretid)
	zone := findZone(aws.Zones(),c.domain)
	resourceRecordSets, err := zone.ResourceRecordSets(aws)

	if err != nil {
		log.Fatal("Resource Record Sets Invalid:", resourceRecordSets, err)
	}

	record := findRecord(resourceRecordSets, c.subdomain, c.domain)

	if record.Name == "" {
		updateRecord(zone, aws, "CREATE", c.subdomain+"."+c.domain, c.ip)
		fmt.Println("Created A record with name ", c.subdomain)
		os.Exit(1)
	}

	fmt.Println("IP was " + record.Value[0])

	if len(record.Value[0]) > 0 {
		fmt.Println("Nothing to do")
		os.Exit(1)
	}

	fmt.Println("Updating IP with Route53")
	updateRecord(zone, aws, "DELETE", c.subdomain+"."+c.domain, record.Value[0])
	updateRecord(zone, aws, "CREATE", c.subdomain+"."+c.domain, c.ip)
	fmt.Println("Done")
	return nil
}

func (c *NewSubdomain) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("dnsass", gnuflag.ExitOnError)
		c.fs.StringVar(&c.accesskey, "accesskey", "", "accesskey: AWS Accesskey")
		c.fs.StringVar(&c.accesskey, "a", "", "accesskey: AWS Accesskey")
		c.fs.StringVar(&c.secretid, "secretid", "", "secretid: AWS Secretid")
		c.fs.StringVar(&c.secretid, "s", "", "secretid: AWS Secretid")
		c.fs.StringVar(&c.subdomain, "subdomain", "", "subdomain: subdomain name")
		c.fs.StringVar(&c.subdomain, "u", "", "subdomain: subdomain name")
		c.fs.StringVar(&c.domain, "domain", "megam.co", "domain: domain name, this needs to preexist in the DNS service. Default : megam.co")
		c.fs.StringVar(&c.domain, "d", "megam.co", "domain: domain name, this needs to preexist in the DNS service")
		c.fs.StringVar(&c.domain, "ipaddress", "", "ipaddress: ipaddress of the running server")
		c.fs.StringVar(&c.domain, "i", "", "ipaddress: ipaddress of the running server")
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
		Usage:   `delete -a <accesskey> -s <secretid> -d <domain name, default:megam.co> -s <subdomain>`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (c *DeleteSubdomain) Run(context *cmd.Context) error {
	return nil
}

func (c *DeleteSubdomain) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("dnsass", gnuflag.ExitOnError)
		c.fs.StringVar(&c.accesskey, "accesskey", "", "accesskey: AWS Accesskey")
		c.fs.StringVar(&c.accesskey, "a", "", "accesskey: AWS Accesskey")
		c.fs.StringVar(&c.secretid, "secretid", "", "secretid: AWS Secretid")
		c.fs.StringVar(&c.secretid, "s", "", "secretid: AWS Secretid")
		c.fs.StringVar(&c.subdomain, "subdomain", "", "subdomain: subdomain name")
		c.fs.StringVar(&c.subdomain, "u", "", "subdomain: subdomain name")
		c.fs.StringVar(&c.domain, "domain", "megam.co", "domain: domain name, this needs to preexist in the DNS service. Default : megam.co")
		c.fs.StringVar(&c.domain, "d", "megam.co", "domain: domain name, this needs to preexist in the DNS service")
	}
	return c.fs
}

func auth(accesskey string, secretkey string) (route53.AccessIdentifiers) {
	return route53.AccessIdentifiers{
		AccessKey: accesskey,
		SecretKey: secretkey,
	}
}

func findZone(zones route53.HostedZones, hosted_zone string) (hz route53.HostedZone) {
	for i := range zones.HostedZones {
		if zones.HostedZones[i].Name == hosted_zone {
			return zones.HostedZones[i]
		}
	}
	return hz
}

func findRecord(records route53.ResourceRecordSets, subdomain string, hosted_zone string) (rrs route53.ResourceRecordSet) {
	for i := range records.ResourceRecordSets {
		if records.ResourceRecordSets[i].Name == subdomain+"."+hosted_zone {
			return records.ResourceRecordSets[i]
		}
	}
	return rrs
}

func updateRecord(zone route53.HostedZone, aws route53.AccessIdentifiers, action string, name string, value string) {
	var create = route53.ChangeResourceRecordSetsRequest{
		ZoneID:  zone.HostedZoneId(),
		Comment: "",
		Changes: []route53.Change{
			{
				Action: action,
				Name:   name,
				Type:   "A",
				TTL:    300,
				Value:  value,
			},
		},
	}

	r, err := create.Create(aws)

	if err != nil {
		log.Fatal("Update record failed:", r, err)
	}
}
