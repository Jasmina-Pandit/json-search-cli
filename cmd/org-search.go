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
)

type OrgSearch struct {
	orgs            []model.Organisation
	orgKeys         []string
	keysOfTypeArray []string
}

func NewOrgSearch(fileName string) *OrgSearch {
	var keysOfTypeArray []string
	keysOfTypeArray = append(append(keysOfTypeArray, "tags"), "domainnames")
	return &OrgSearch{
		orgs:            loadOrgs(fileName),
		orgKeys:         structs.Names(&model.Organisation{}),
		keysOfTypeArray: keysOfTypeArray,
	}
}
func loadOrgs(fileName string) []model.Organisation {
	jsonReader := reader.NewReader()
	var orgs []model.Organisation
	byte, _ := jsonReader.ReadFixtureFile(fileName)
	err := json.Unmarshal(byte, &orgs)
	if err != nil {
		fmt.Print(err)
	}
	return orgs
}
func (*OrgSearch) Help() string {
	return "hello"
}

func (o *OrgSearch) searchOrg(key string, value string) ([]model.Organisation, error) {
	var result []model.Organisation
	if !helper.IsCaseAndUnderscoreInsenKeyInArray(o.orgKeys, key) {
		return nil, errors.New("invalid key. Use help command for list of valid keys")
	}

	for _, org := range o.orgs {
		u := reflect.ValueOf(org)
		v := helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if fmt.Sprint(v) == value || helper.CheckTrimmedValueInArrayString(key, o.keysOfTypeArray, v, value) {
			o.printPretty(org)
			result = append(result, org)
		}

	}
	return result, nil
}

func (o *OrgSearch) printPretty(orgs model.Organisation) {
	org := reflect.ValueOf(orgs)
	fmt.Println("ORGANISATIONS")
	for _, key := range o.orgKeys {
		fmt.Println(key, ":", helper.CaseAndUnderscoreInsenstiveFieldByName(org, key))
	}
}

func (u *OrgSearch) Run(args []string) int {
	u.searchOrg(args[0], args[1])
	return 0
}
func (h *OrgSearch) Synopsis() string {
	return h.Help()
}
