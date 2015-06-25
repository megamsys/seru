package main

import (
	"fmt"
	"github.com/megamsys/seru/cmd"
	"github.com/karlentwistle/route53"
	"launchpad.net/gnuflag"
	"log"
	"os"
)

type Listdomain struct {
	fs        *gnuflag.FlagSet
	accesskey string
	secretid  string
}

func (l *Listdomain) Info() *cmd.Info {
	desc := `lists the domains in the DNS service.`
	return &cmd.Info{
		Name:    "list",
		Usage:   `list -a <accesskey> -s <secretid>`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (l *Listdomain) Run(context *cmd.Context) error {
	aws := auth(l.accesskey, l.secretid)
	zones := aws.Zones().HostedZones
	
	table := cmd.NewTable()
	table.Headers = cmd.Row([]string{"Id", "Name"})
	for _, zone := range zones {		
		table.AddRow(cmd.Row([]string{zone.Id, zone.Name}))
	}
	table.Sort()
	context.Stdout.Write(table.Bytes())	
	return nil
}

func (l *Listdomain) Flags() *gnuflag.FlagSet {
	if l.fs == nil {
		l.fs = gnuflag.NewFlagSet("dnsasslist", gnuflag.ExitOnError)
		l.fs.StringVar(&l.accesskey, "accesskey", "", "accesskey: AWS Accesskey")
		l.fs.StringVar(&l.accesskey, "a", "", "accesskey: AWS Accesskey")
		l.fs.StringVar(&l.secretid, "secretid", "", "secretid: AWS Secretid")
		l.fs.StringVar(&l.secretid, "s", "", "secretid: AWS Secretid")
	}
	return l.fs
}

type NewSubdomain struct {
	Fs        *gnuflag.FlagSet
	Accesskey string
	Secretid  string
	Domain    string
	Subdomain string
	Ip        string
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
	aws := auth(c.Accesskey, c.Secretid)
	zone := findZone(aws.Zones(), c.Domain)
		
	resourceRecordSets, err := zone.ResourceRecordSets(aws)

	if err != nil {
		log.Fatal("Resource Record Sets Invalid:", resourceRecordSets, err)
	}

	record := findRecord(resourceRecordSets, c.Subdomain, c.Domain)

	if record.Name == "" {
		updateRecord(zone, aws, "CREATE", c.Subdomain+"."+c.Domain, c.Ip)
		fmt.Println("Created A record with name ", c.Subdomain)
		os.Exit(0)
	}

	fmt.Println("IP was " + record.Value[0])

	if len(record.Value[0]) > 0 {
		fmt.Println("Nothing to do")
		os.Exit(0)
	}

	fmt.Println("Updating IP with Route53")
	updateRecord(zone, aws, "DELETE", c.Subdomain+"."+c.Domain, record.Value[0])
	updateRecord(zone, aws, "CREATE", c.Subdomain+"."+c.Domain, c.Ip)
	fmt.Println("Done")
	return nil
}

func (c *NewSubdomain) ApiRun(context *cmd.Context) error {
	aws := auth(c.Accesskey, c.Secretid)
	zone := findZone(aws.Zones(), c.Domain)
		
	resourceRecordSets, err := zone.ResourceRecordSets(aws)

	if err != nil {
		log.Fatal("Resource Record Sets Invalid:", resourceRecordSets, err)
	}

	record := findRecord(resourceRecordSets, c.Subdomain, c.Domain)

	if record.Name == "" {
		err := updateRecord(zone, aws, "CREATE", c.Subdomain+"."+c.Domain, c.Ip)
		if err != nil {
	   		return err
		}
		fmt.Println("Created A record with name ", c.Subdomain)
		return nil
	}

	fmt.Println("IP was " + record.Value[0])

	if len(record.Value[0]) > 0 {
		fmt.Println("Nothing to do")
		return nil
	}

	fmt.Println("Updating IP with Route53")
	del_err := updateRecord(zone, aws, "DELETE", c.Subdomain+"."+c.Domain, record.Value[0])
	if del_err != nil {
	   return del_err
	}
	crt_err := updateRecord(zone, aws, "CREATE", c.Subdomain+"."+c.Domain, c.Ip)
	if crt_err != nil {
	   return crt_err
	}
	fmt.Println("Done")
	return nil
}

func (c *NewSubdomain) Flags() *gnuflag.FlagSet {
	if c.Fs == nil {
		c.Fs = gnuflag.NewFlagSet("dnsassnew", gnuflag.ExitOnError)
		c.Fs.StringVar(&c.Accesskey, "accesskey", "", "accesskey: AWS Accesskey")
		c.Fs.StringVar(&c.Accesskey, "a", "", "accesskey: AWS Accesskey")
		c.Fs.StringVar(&c.Secretid, "secretid", "", "secretid: AWS Secretid")
		c.Fs.StringVar(&c.Secretid, "s", "", "secretid: AWS Secretid")
		c.Fs.StringVar(&c.Subdomain, "subdomain", "", "subdomain: subdomain name")
		c.Fs.StringVar(&c.Subdomain, "u", "", "subdomain: subdomain name")
		c.Fs.StringVar(&c.Domain, "domain", "megam.co", "domain: domain name, this needs to preexist in the DNS service. Default : megam.co")
		c.Fs.StringVar(&c.Domain, "d", "megam.co", "domain: domain name, this needs to preexist in the DNS service")
		c.Fs.StringVar(&c.Ip, "ipaddress", "", "ipaddress: ipaddress of the running server")
		c.Fs.StringVar(&c.Ip, "i", "", "ipaddress: ipaddress of the running server")
	}
	return c.Fs
}

type DeleteSubdomain struct {
	fs        *gnuflag.FlagSet
	yes       bool
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
	aws := auth(c.accesskey, c.secretid)
	zone := findZone(aws.Zones(), c.domain)
	resourceRecordSets, err := zone.ResourceRecordSets(aws)

	if err != nil {
		log.Fatal("Resource Record Sets Invalid:", resourceRecordSets, err)
	}

	record := findRecord(resourceRecordSets, c.subdomain, c.domain)

	if record.Name == "" {
		fmt.Println("A record not found with name ", c.subdomain+"."+c.domain)
		os.Exit(1)
	}

	var answer string
	if !c.yes {
		fmt.Fprintf(context.Stdout, `Are you sure you want to remove "%s"? (y/n) `, c.subdomain+"."+c.domain)
		fmt.Fscanf(context.Stdin, "%s", &answer)
		if answer != "y" {
			fmt.Fprintln(context.Stdout, "Abort.")
			return nil
		}
	}
	updateRecord(zone, aws, "DELETE", c.subdomain+"."+c.domain, record.Value[0])
	fmt.Println("Done")
	return nil
}

func (c *DeleteSubdomain) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("dnsassdel", gnuflag.ExitOnError)
		c.fs.StringVar(&c.accesskey, "accesskey", "", "accesskey: AWS Accesskey")
		c.fs.StringVar(&c.accesskey, "a", "", "accesskey: AWS Accesskey")
		c.fs.StringVar(&c.secretid, "secretid", "", "secretid: AWS Secretid")
		c.fs.StringVar(&c.secretid, "s", "", "secretid: AWS Secretid")
		c.fs.StringVar(&c.subdomain, "subdomain", "", "subdomain: subdomain name")
		c.fs.StringVar(&c.subdomain, "u", "", "subdomain: subdomain name")
		c.fs.StringVar(&c.domain, "domain", "megam.co", "domain: domain name, this needs to preexist in the DNS service. Default : megam.co")
		c.fs.StringVar(&c.domain, "d", "megam.co", "domain: domain name, this needs to preexist in the DNS service")
		c.fs.BoolVar(&c.yes, "assume-yes", false, "Don't ask for confirmation, just remove the subdomain.")
		c.fs.BoolVar(&c.yes, "y", false, "Don't ask for confirmation, just remove the subdomain.")
	}
	return c.fs
}

func auth(accesskey string, secretkey string) route53.AccessIdentifiers {
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

func updateRecord(zone route53.HostedZone, aws route53.AccessIdentifiers, action string, name string, value string) error {
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
		return err
	}
	
	return nil
}
