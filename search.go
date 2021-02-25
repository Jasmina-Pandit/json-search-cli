package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

func main() {

	c:= cli.NewCLI("json-search-cli","0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"search-user": func() (cli.Command, error) {
			return &UserSearch{}, nil
		},
	}
	status, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}

type UserSearch struct {

}

func (*UserSearch) Help() string {
	return "hello"
}

func (u *UserSearch) Run(args []string) int{
	fmt.Print("Hello")
	return 0
}
func (h *UserSearch) Synopsis() string {
	return h.Help()
}
