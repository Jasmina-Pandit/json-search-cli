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

type OrgSearch struct {
	orgs []model.Organisation
}
func NewOrgSearch() *OrgSearch{
	return &OrgSearch{orgs: loadOrgs()}
}
func loadOrgs() []model.Organisation {
	jsonReader:= reader.NewReader()
	var orgs []model.Organisation
	byte, _:= jsonReader.ReadFixtureFile("ref-data/organizations.json")
	err:= json.Unmarshal(byte,&orgs)
	if err !=nil{
		fmt.Print(err)
	}
	return orgs
}
func (*OrgSearch) Help() string {
	return "hello"
}

func (o *OrgSearch) searchOrg(key string, value string){

	for _, org := range o.orgs {
		u := reflect.ValueOf(org)
		v:= helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if fmt.Sprint(v) == value{
			o.printPretty(org)
		}

	}
}

func (o *OrgSearch) printPretty(orgs model.Organisation){
	org := reflect.ValueOf(orgs)
	fmt.Println("ORGANISATIONS")
	for _,key:= range structs.Names(&model.Organisation{}){
		fmt.Println(key,":",helper.CaseAndUnderscoreInsenstiveFieldByName(org,key))
	}
}

func (u *OrgSearch) Run(args []string) int{
	u.searchOrg(args[0], args[1])
	return 0
}
func (h *OrgSearch) Synopsis() string {
	return h.Help()
}