package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/sirupsen/logrus"
	"github.com/zendesk/json-search-cli/model"
	"github.com/zendesk/json-search-cli/reader"
	"os"
)

func main() {

	c := cli.NewCLI("json-search-cli", "0.0.1")
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

func (u *UserSearch) Run(args []string) int {
	users := loadUser()
	logrus.Info("Hello", users)
	return 0
}
func (h *UserSearch) Synopsis() string {
	return h.Help()
}

func loadUser() []model.User {
	jsonReader := reader.NewReader()
	var users []model.User
	byte, _ := jsonReader.ReadFixtureFile("ref-data/users.json")
	err := json.Unmarshal(byte, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}
