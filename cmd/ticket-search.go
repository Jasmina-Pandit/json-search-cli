package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"json-search-cli/helper"
	"json-search-cli/model"
	"json-search-cli/reader"
	"reflect"
)

type TicketSearch struct {
	tickets []model.Ticket
}
func NewTicketSearch() *TicketSearch{
	return &TicketSearch{tickets: loadTickets()}
}
func loadTickets() []model.Ticket {
	jsonReader:= reader.NewReader()
	var tkts []model.Ticket
	byte, _:= jsonReader.ReadFixtureFile("ref-data/tickets.json")
	err:= json.Unmarshal(byte,&tkts)
	if err !=nil{
		fmt.Print(err)
	}
	return tkts
}
func (*TicketSearch) Help() string {
	return "hello"
}

func (t *TicketSearch) searchTicket(key string, value string){

	for _, tkt := range t.tickets {
		u := reflect.ValueOf(tkt)
		v:= helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if fmt.Sprint(v) == value{
			t.printPretty(tkt)
		}
	}
}

func (t *TicketSearch) printPretty(tkts model.Ticket){
	tkt := reflect.ValueOf(tkts)
	fmt.Println("TICKETS")
	for _,key:= range structs.Names(&model.Ticket{}){
		fmt.Println(key,":",helper.CaseAndUnderscoreInsenstiveFieldByName(tkt,key))
	}
}

func (t *TicketSearch) Run(args []string) int{
	t.searchTicket(args[0], args[1])
	return 0
}
func (h *TicketSearch) Synopsis() string {
	return h.Help()
}