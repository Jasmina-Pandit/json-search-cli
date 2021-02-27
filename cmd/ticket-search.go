package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"json-search-cli/helper"
	"json-search-cli/model"
	"json-search-cli/reader"
	"reflect"
	"strings"
)

type TicketSearch struct {
	tickets         []model.Ticket
	tktKeys         []string
	keysOfTypeArray []string
}

func NewTicketSearch(fileName string) *TicketSearch {
	var keysOfTypeArray []string
	keysOfTypeArray = append(keysOfTypeArray, "tags")

	return &TicketSearch{
		tickets:         loadTickets(fileName),
		tktKeys:         structs.Names(&model.Ticket{}),
		keysOfTypeArray: keysOfTypeArray,
	}
}
func loadTickets(fileName string) []model.Ticket {
	jsonReader := reader.NewReader()
	var tkts []model.Ticket
	byte, _ := jsonReader.ReadFixtureFile(fileName)
	err := json.Unmarshal(byte, &tkts)
	if err != nil {
		fmt.Print(err)
	}
	return tkts
}
func (*TicketSearch) Help() string {
	return "hello"
}

func (t *TicketSearch) searchTicket(key string, value string) ([]model.Ticket, error) {
	var result []model.Ticket
	if !helper.IsCaseAndUnderscoreInsenKeyInArray(t.tktKeys, key) {
		return nil, errors.New("invalid key. Use help command for list of valid keys")
	}
	for _, tkt := range t.tickets {
		u := reflect.ValueOf(tkt)
		v := helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if strings.ToLower(fmt.Sprint(v)) == strings.ToLower(value) || helper.CheckTrimmedValueInArrayString(key, t.keysOfTypeArray, v, value) {
			t.printPretty(tkt)
			result = append(result, tkt)
		}
	}
	return result, nil
}

func (t *TicketSearch) printPretty(tkts model.Ticket) {
	tkt := reflect.ValueOf(tkts)
	fmt.Println("TICKETS")
	for _, key := range t.tktKeys {
		fmt.Println(key, ":", helper.CaseAndUnderscoreInsenstiveFieldByName(tkt, key))
	}
}

func (t *TicketSearch) Run(args []string) int {
	t.searchTicket(args[0], args[1])
	return 0
}
func (h *TicketSearch) Synopsis() string {
	return h.Help()
}
