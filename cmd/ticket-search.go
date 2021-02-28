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
	"strconv"
	"strings"
)

type TicketSearch struct {
	tickets         []model.Ticket
	tktKeys         []string
	keysOfTypeArray []string
	userSearch      *UserSearch
	orgSearch       *OrgSearch
	userFileName    string
	orgFileName     string
	tktFileName     string
}

func NewTicketSearch(tktFileName string, userFileName string, orgFileName string) *TicketSearch {

	return &TicketSearch{
		tktKeys:      structs.Names(&model.Ticket{}),
		userFileName: userFileName,
		orgFileName:  orgFileName,
		tktFileName:  tktFileName,
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

func (t *TicketSearch) searchTicket(key string, value string) (*model.Response, error) {
	var result []model.Ticket
	if !helper.IsCaseAndUnderscoreInsenKeyInArray(t.tktKeys, key) {
		return nil, errors.New("invalid key. Use help command for list of valid keys")
	}
	userResult := []model.User{}
	orgResult := []model.Organisation{}
	for _, tkt := range t.tickets {
		u := reflect.ValueOf(tkt)
		v := helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if strings.ToLower(fmt.Sprint(v)) == strings.ToLower(value) || helper.CheckTrimmedValueInArrayString(key, t.keysOfTypeArray, v, value) {
			t.printPretty(tkt)
			userResp, err := t.userSearch.searchUser("ID", strconv.Itoa(tkt.AssigneeID))
			if err != nil {
				logrus.WithField("Ticket ID", tkt.ID).WithField("User ID", tkt.AssigneeID).WithError(err).Error("Failed to fetch assigned user for ticket")
			} else if userResp == nil || len(userResp.Users) <= 0 {
				logrus.WithField("Ticket ID", tkt.ID).WithField("User ID", tkt.AssigneeID).Warn("No corresponding user found for the ticket")
			} else {
				userResult = append(userResult, userResp.Users...)
			}

			orgResp, err := t.orgSearch.searchOrg("ID", strconv.Itoa(tkt.OrganizationID))
			if err != nil {
				logrus.WithField("Ticket ID", tkt.ID).WithField("Org ID", tkt.OrganizationID).WithError(err).Error("Failed to fetch organisation for ticket")
			} else if orgResp == nil || len(orgResp.Orgs) <= 0 {
				logrus.WithField("Ticket ID", tkt.ID).WithField("Org ID", tkt.OrganizationID).Warn("No corresponding organisation found for the ticket")
			} else {
				orgResult = append(orgResult, orgResp.Orgs...)
			}
			result = append(result, tkt)
		}
	}
	response := model.Response{Users: userResult, Orgs: orgResult, Tickets: result}
	return &response, nil
}

func (t *TicketSearch) printPretty(tktModel model.Ticket) {
	tkt := reflect.ValueOf(tktModel)
	fmt.Println("TICKETS")

	for _, key := range t.tktKeys {
		fmt.Println(key, ":", helper.CaseAndUnderscoreInsenstiveFieldByName(tkt, key))
	}
}

func (t *TicketSearch) Run(args []string) int {
	if len(args) < 2 {
		return -18511 //return help as per cli doco
	}
	err := t.Initialise()
	if err != nil {
		return -1
	}
	t.searchTicket(args[0], args[1])

	return 0
}
func (t *TicketSearch) Initialise() error {
	t.userSearch = NewUserSearch(t.userFileName, t.orgFileName)
	t.userSearch.Initialise()
	if len(t.userSearch.users) == 0 {
		logrus.Error("No users found")
		return errors.New("no user found. check reference users json file")
	}
	t.orgSearch = NewOrgSearch(t.orgFileName)
	t.orgSearch.Initialise()
	if len(t.orgSearch.orgs) == 0 {
		logrus.Error("No organisations found")
		return errors.New("no organisations found. check reference organisation json file")
	}
	t.tickets = loadTickets(t.tktFileName)
	if len(t.tickets) == 0 {
		logrus.Error("No tickets found")
		return errors.New("no tickets found. check reference tickets json file")
	}
	t.keysOfTypeArray = append(t.keysOfTypeArray, "tags")
	return nil
}

func (t *TicketSearch) Synopsis() string {
	return t.Help()
}

func (*TicketSearch) Help() string {
	return "\tSearch Ticket and it assigned user and organisation using the search key and field. Key and value are case and underscore agnostic \n" +
		"\t\t\tSyntax search-tkt <key> <value>  \n" +
		"\t\t\te.g: search-tkt id 1"
}
