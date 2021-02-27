package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"json-search-cli/cmd"
	"os"
)

const (
	orgFileName    = "ref-data/organizations.json"
	userFileName   = "ref-data/users.json"
	ticketFileName = "ref-data/organizations.json"
)

func main() {

	c := cli.NewCLI("json-search-cli", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"search-user": func() (cli.Command, error) {
			return cmd.NewUserSearch(userFileName, orgFileName), nil
		},
		"search-org": func() (cli.Command, error) {
			return cmd.NewOrgSearch(orgFileName), nil
		},
		"search-tkt": func() (cli.Command, error) {
			return cmd.NewTicketSearch(ticketFileName, userFileName, orgFileName), nil
		},
	}
	status, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}
