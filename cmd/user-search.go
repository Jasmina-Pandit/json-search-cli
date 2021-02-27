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

type UserSearch struct {
	users           []model.User
	userKeys        []string
	keysOfTypeArray []string
	orgSearch       *OrgSearch
}

func NewUserSearch(userFileName string, orgFileName string) *UserSearch {
	var keysOfTypeArray []string
	keysOfTypeArray = append(keysOfTypeArray, "tags")
	orgSearch := NewOrgSearch(orgFileName)
	return &UserSearch{
		users:           loadUser(userFileName),
		userKeys:        structs.Names(&model.User{}),
		keysOfTypeArray: keysOfTypeArray,
		orgSearch:       orgSearch,
	}
}
func loadUser(fileName string) []model.User {
	jsonReader := reader.NewReader()
	var users []model.User
	byte, _ := jsonReader.ReadFixtureFile(fileName)
	err := json.Unmarshal(byte, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}

func (*UserSearch) Run1(args []string) int {
	users := loadUser("ref-data/users.json")
	userKeys := structs.Names(&model.User{})
	userMap := make(map[string]map[string][]model.User)

	for _, user := range users {
		u := reflect.ValueOf(user)

		for _, key := range userKeys {
			if _, found := userMap[key]; !found {
				valueMap := make(map[string][]model.User)
				var userSlice []model.User
				userSlice = append(userSlice, user)
				valueMap[(u.FieldByName(key).Interface().(string))] = userSlice
				userMap[key] = valueMap
			} else {
				valueMap := userMap[key]
				valueKey := u.FieldByName(key).Interface().(string)
				if _, found := valueMap[valueKey]; !found {
					userSlice := valueMap[valueKey]
					userSlice = append(userSlice, user)
					valueMap[(u.FieldByName(key).Interface().(string))] = userSlice
					userMap[key] = valueMap
				} else {
					var userSlice []model.User
					userSlice = append(userSlice, user)
					valueMap[(u.FieldByName(key).Interface().(string))] = userSlice
					userMap[key] = valueMap
				}
			}

		}

	}
	fmt.Println(userMap)
	return 0
}

//search user by key and value in user.json
func (u *UserSearch) searchUser(key string, value string) (*model.Response, error) {

	//validate key
	if !helper.IsCaseAndUnderscoreInsenKeyInArray(u.userKeys, key) {
		return nil, errors.New("invalid key. Use help command for list of valid keys")
	}
	userResult := []model.User{}
	orgResult := []model.Organisation{}

	for _, user := range u.users {
		ru := reflect.ValueOf(user)
		v := helper.CaseAndUnderscoreInsenstiveFieldByName(ru, key)

		if strings.ToLower(fmt.Sprint(v)) == strings.ToLower(value) || helper.CheckTrimmedValueInArrayString(key, u.keysOfTypeArray, v, value) {
			u.printPretty(user)
			response, err := u.orgSearch.searchOrg("ID", strconv.Itoa(user.OrganizationID))
			if err != nil {
				logrus.WithField("User ID", user.Id).WithField("Org ID", user.OrganizationID).WithError(err).Error("Failed to fetch organisation for user")
			} else if response == nil || len(response.Orgs) == 0 {
				logrus.WithField("User ID", user.Id).WithField("Org ID", user.OrganizationID).Warn("No organisation found for user")
			} else {
				orgResult = append(orgResult, response.Orgs...)
			}

			userResult = append(userResult, user)
		}
	}
	response := model.Response{Users: userResult, Orgs: orgResult}
	return &response, nil
}

func (u *UserSearch) printPretty(user model.User) {
	ru := reflect.ValueOf(user)
	fmt.Println("USER")

	for _, key := range u.userKeys {
		fmt.Println(key, ":", helper.CaseAndUnderscoreInsenstiveFieldByName(ru, key))
	}
}

func (u *UserSearch) Run(args []string) int {
	u.searchUser(args[0], args[1])
	return 0
}
func (h *UserSearch) Synopsis() string {
	return h.Help()
}

func (*UserSearch) Help() string {
	return "\tSearch User and its organisation using the search key and field. Key and value are case and underscore agnostic \n" +
		"\t\t\tSyntax search-user <key> <value>  \n" +
		"\t\t\te.g: search-tkt id 1"
}
