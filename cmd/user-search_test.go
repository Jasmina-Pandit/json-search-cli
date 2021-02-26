package cmd

import (
	"github.com/stretchr/testify/require"
	"json-search-cli/model"
	"testing"
)

func Test_UserSearch(t *testing.T) {

	tests := map[string]struct {
		thenAssert       func(users []model.User, err error)
		givenKey         string
		givenSearchValue string
	}{
		"should successfully search small case id": {
			givenKey:         "id",
			givenSearchValue: "1",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
			},
		},
		"should successfully search upper case id": {
			givenKey:         "ID",
			givenSearchValue: "1",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
			},
		},
		"should successfully search underscore id": {
			givenKey:         "_id",
			givenSearchValue: "1",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
			},
		},
		"should successfully search by external_id": {
			givenKey:         "external_id",
			givenSearchValue: "12ab34",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
			},
		},
		"should successfully search by externalid without underscore": {
			givenKey:         "externalid",
			givenSearchValue: "12ab34",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 1)
			},
		},
		"should successfully search for value in array fields": {
			givenKey:         "tags",
			givenSearchValue: "Springville",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 2)
			},
		},
		"should successfully search for value with spaces": {
			givenKey:         "signature",
			givenSearchValue: "Don't Worry Be Happy!",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.NotNil(t, users)
				require.Len(t, users, 2)
			},
		},
		"should not error on missing fields": {
			givenKey:         "newkey",
			givenSearchValue: "hello",
			thenAssert: func(users []model.User, err error) {
				require.NotNil(t, err)
				require.Error(t, err, "invalid key. Use help command for list of valid keys")

			},
		},
		"should return user for empty searches": {
			givenKey:         "alias",
			givenSearchValue: "",
			thenAssert: func(users []model.User, err error) {
				require.Nil(t, err)
				require.Len(t, users, 1)
			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userSearch := NewUserSearch("testdata/user_test.json")
			result, err := userSearch.searchUser(test.givenKey, test.givenSearchValue)
			test.thenAssert(result, err)
		})
	}
}
