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

type UserSearch struct {
	users []model.User
}
func NewUserSearch() *UserSearch{
	return &UserSearch{users: loadUser()}
}
func loadUser() []model.User {
	jsonReader:= reader.NewReader()
	var users []model.User
	byte, _:= jsonReader.ReadFixtureFile("ref-data/users.json")
	err:= json.Unmarshal(byte,&users)
	if err !=nil{
		fmt.Print(err)
	}
	return users
}
func (*UserSearch) Help() string {
	return "hello"
}
func (*UserSearch) Run1(args []string) int {
	users:=loadUser()
	userKeys:= structs.Names(&model.User{})
	userMap := make(map[string]map[string][]model.User)

	for _,user:=range users{
		u:= reflect.ValueOf(user)

		for _, key:= range userKeys{
			if _, found := userMap[key]; !found {
				valueMap:= make(map[string][]model.User)
				var userSlice []model.User
				userSlice=append(userSlice, user)
				valueMap[(u.FieldByName(key).Interface().(string))]=userSlice
				userMap[key]=valueMap
			}else{
				valueMap:= userMap[key]
				valueKey:= u.FieldByName(key).Interface().(string)
				if _, found := valueMap[valueKey]; !found {
					userSlice:= valueMap[valueKey]
					userSlice= append(userSlice, user)
					valueMap[(u.FieldByName(key).Interface().(string))] =userSlice
					userMap[key]=valueMap
				}else {
					var userSlice []model.User
					userSlice=append(userSlice, user)
					valueMap[(u.FieldByName(key).Interface().(string))]=userSlice
					userMap[key]=valueMap
				}
			}

		}

	}
	fmt.Println(userMap)
	return 0
}

func (u *UserSearch) searchUser(key string, value string){
	users:=loadUser()
	for _,user:= range users {
		u := reflect.ValueOf(user)
		v:= helper.CaseAndUnderscoreInsenstiveFieldByName(u, key)
		if fmt.Sprint(v) == value{
			printPretty(user)
		}

	}
}

func printPretty(user model.User){
	u := reflect.ValueOf(user)
	fmt.Println("USER")
	for _,key:= range structs.Names(&model.User{}){
		fmt.Println(key,":",helper.CaseAndUnderscoreInsenstiveFieldByName(u,key))
	}
}

func (u *UserSearch) Run(args []string) int{
	u.searchUser(args[0], args[1])
	return 0
}
func (h *UserSearch) Synopsis() string {
	return h.Help()
}