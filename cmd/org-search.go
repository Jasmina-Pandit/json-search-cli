package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
	"json-search-cli/helper"
	"json-search-cli/model"
	"json-search-cli/reader"
	"reflect"
	"strings"
)

type OrgSearch struct {
	orgs            []model.Organisation
	orgKeys         []string
	keysOfTypeArray []string
	orgfileName     string
}

func NewOrgSearch(fileName string) *OrgSearch {
	return &OrgSearch{
		orgfileName: fileName,
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

func (o *OrgSearch) searchOrg(key string, value string) (*model.Response, error) {
	var result []model.Organisation
	if !helper.IsCaseAndUnderscoreInsenKeyInArray(o.orgKeys, key) {
		return nil, errors.New("invalid key. Use help command for list of valid keys")
	}

	for _, org := range o.orgs {
		u := reflect.ValueOf(org)
		v := helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if strings.ToLower(fmt.Sprint(v)) == strings.ToLower(value) || helper.CheckTrimmedValueInArrayString(key, o.keysOfTypeArray, v, value) {
			o.printPretty(org)
			result = append(result, org)
		}

	}
	response := model.Response{Orgs: result}
	return &response, nil
}

func (o *OrgSearch) printPretty(orgs model.Organisation) {
	org := reflect.ValueOf(orgs)
	fmt.Println("ORGANISATIONS")
	for _, key := range o.orgKeys {
		fmt.Println(key, ":", helper.CaseAndUnderscoreInsenstiveFieldByName(org, key))
	}
}

func (o *OrgSearch) Run(args []string) int {
	if len(args) < 2 {
		return -18511 //return help as per cli doco
	}
	err := o.Initialise()
	if err != nil {
		return -1
	}
	o.searchOrg(args[0], args[1])

	return 0
}

func (o *OrgSearch) Initialise() error {
	o.orgs = loadOrgs(o.orgfileName)
	if len(o.orgs) == 0 {
		logrus.Error("No organisations found")
		return errors.New("no organisations found. check reference organisation json file")
	}
	o.orgKeys = structs.Names(&model.Organisation{})
	o.keysOfTypeArray = append(append(o.keysOfTypeArray, "tags"), "domainnames")
	return nil
}

func (o *OrgSearch) Synopsis() string {
	return o.Help()
}

func (*OrgSearch) Help() string {
	return "\tSearch Organisation using the search key and field. Key and value are case and underscore agnostic \n" +
		"\t\t\tSyntax search-org <key> <value>  \n" +
		"\t\t\te.g: search-org id 1"
}
