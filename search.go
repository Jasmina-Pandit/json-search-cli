package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"json-search-cli/cmd"
	"os"
)

func main() {

	c := cli.NewCLI("json-search-cli", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"search-user": func() (cli.Command, error) {
			return cmd.NewUserSearch("ref-data/users.json"), nil
		},
		"search-org": func() (cli.Command, error) {
			return cmd.NewOrgSearch("ref-data/organizations.json"), nil
		},
		"search-tkt": func() (cli.Command, error) {
			return cmd.NewTicketSearch("ref-data/organizations.json"), nil
		},
	}
	status, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}
