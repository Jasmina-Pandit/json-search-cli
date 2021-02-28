package cmd

import (
	"github.com/stretchr/testify/require"
	"json-search-cli/model"
	"testing"
)

func Test_UserSearch(t *testing.T) {

	tests := map[string]struct {
		thenAssert       func(response *model.Response, err error)
		givenKey         string
		givenSearchValue string
	}{
		"should successfully search small case id": {
			givenKey:         "id",
			givenSearchValue: "1",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
				require.Equal(t, users[0].Id, 1)
			},
		},
		"should successfully search upper case id": {
			givenKey:         "ID",
			givenSearchValue: "1",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
				require.Equal(t, users[0].Id, 1)
			},
		},
		"should successfully search underscore id": {
			givenKey:         "_id",
			givenSearchValue: "1",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
				require.Equal(t, users[0].Id, 1)
			},
		},
		"should successfully search by external_id": {
			givenKey:         "external_id",
			givenSearchValue: "12ab34",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
				require.Equal(t, users[0].ExternalID, "12ab34")
			},
		},
		"should successfully search by externalid without underscore": {
			givenKey:         "externalid",
			givenSearchValue: "12ab34",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
				require.Equal(t, users[0].ExternalID, "12ab34")
			},
		},
		"should successfully search for value in array fields": {
			givenKey:         "tags",
			givenSearchValue: "Springville",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 2)
			},
		},
		"should successfully search for value with spaces": {
			givenKey:         "signature",
			givenSearchValue: "Don't Worry Be Happy!",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 2)
			},
		},
		"should not error on missing fields": {
			givenKey:         "newkey",
			givenSearchValue: "hello",
			thenAssert: func(response *model.Response, err error) {
				require.NotNil(t, err)
				require.Error(t, err, "invalid key. Use help command for list of valid keys")

			},
		},
		"should return user for empty searches": {
			givenKey:         "alias",
			givenSearchValue: "",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.Len(t, users, 1)
			},
		},
		"should return user for case insensitive searches - all lower": {
			givenKey:         "locale",
			givenSearchValue: "zh-cn",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.Len(t, users, 1)
				require.Equal(t, users[0].Locale, "zh-CN")
			},
		},
		"should return user for case insensitive searches - all upper": {
			givenKey:         "locale",
			givenSearchValue: "ZH-CN",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.Len(t, users, 1)
				require.Equal(t, users[0].Locale, "zh-CN")
			},
		},
		"should return user for case insensitive searches - mixed": {
			givenKey:         "locale",
			givenSearchValue: "Zh-cN",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.Len(t, users, 1)
				require.Equal(t, users[0].Locale, "zh-CN")
			},
		},
		"should return user for case insensitive searches in array fields - all lower": {
			givenKey:         "tags",
			givenSearchValue: "springville",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Nil(t, err)
				require.Len(t, users, 2)
			},
		},
		"should return related organisation for the user": {
			givenKey:         "ID",
			givenSearchValue: "1",
			thenAssert: func(response *model.Response, err error) {
				users := response.Users
				require.Len(t, users, 1)
				require.Equal(t, users[0].Id, 1)
				require.Equal(t, users[0].OrganizationID, 101)

				orgs := response.Orgs
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ID, 101)
			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userSearch := NewUserSearch("testdata/user_test.json", "testdata/org_test.json")
			userSearch.Initialise()
			result, err := userSearch.searchUser(test.givenKey, test.givenSearchValue)
			test.thenAssert(result, err)
		})
	}
}

func Test_UserCommands(t *testing.T) {
	tests := map[string]struct {
		thenAssert        func(result int)
		givenArgs         []string
		givenUserFileName string
		givenOrgFileName  string
	}{
		"should return success for valid key value": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, 0)

			},
		},
		"should return help code when no arguments are passed": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenArgs:         []string{},
			thenAssert: func(result int) {
				require.Equal(t, result, -18511)

			},
		},
		"should return help code when only 1 argument is passed": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenArgs:         []string{},
			thenAssert: func(result int) {
				require.Equal(t, result, -18511)

			},
		},
		"should return success and ignore extra arguments": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, 0)

			},
		},
		"should return -1 for invalid user file": {
			givenUserFileName: "testdata/user.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, -1)

			},
		},
		"should return -1 for invalid org file": {
			givenUserFileName: "testdata/user.json",
			givenOrgFileName:  "testdata/org.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, -1)

			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userSearch := NewUserSearch(test.givenUserFileName, test.givenOrgFileName)
			result := userSearch.Run(test.givenArgs)
			test.thenAssert(result)
		})
	}
}
