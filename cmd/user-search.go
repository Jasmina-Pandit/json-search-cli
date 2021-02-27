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

type UserSearch struct {
	users           []model.User
	userKeys        []string
	keysOfTypeArray []string
}

func NewUserSearch(fileName string) *UserSearch {
	var keysOfTypeArray []string
	keysOfTypeArray = append(keysOfTypeArray, "tags")
	return &UserSearch{
		users:           loadUser(fileName),
		userKeys:        structs.Names(&model.User{}),
		keysOfTypeArray: keysOfTypeArray,
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

func (*UserSearch) Help() string {
	return "hello"
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
func (u *UserSearch) searchUser(key string, value string) ([]model.User, error) {

	//validate key
	if !helper.IsCaseAndUnderscoreInsenKeyInArray(u.userKeys, key) {
		return nil, errors.New("invalid key. Use help command for list of valid keys")
	}

	result := []model.User{}
	for _, user := range u.users {
		ru := reflect.ValueOf(user)
		v := helper.CaseAndUnderscoreInsenstiveFieldByName(ru, key)

		if strings.ToLower(fmt.Sprint(v)) == strings.ToLower(value) || helper.CheckTrimmedValueInArrayString(key, u.keysOfTypeArray, v, value) {
			u.printPretty(user)
			result = append(result, user)
		}
	}
	return result, nil
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
