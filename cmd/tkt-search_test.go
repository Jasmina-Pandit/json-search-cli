package cmd

import (
	"github.com/stretchr/testify/require"
	"json-search-cli/model"
	"testing"
)

func Test_TicketSearch(t *testing.T) {

	tests := map[string]struct {
		thenAssert       func(response *model.Response, err error)
		givenKey         string
		givenSearchValue string
	}{
		"should successfully search small case id": {
			givenKey:         "id",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search upper case id": {
			givenKey:         "ID",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search underscore id": {
			givenKey:         "_id",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search by external_id": {
			givenKey:         "external_id",
			givenSearchValue: "12ab34",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search by externalid without underscore": {
			givenKey:         "externalid",
			givenSearchValue: "12ab34",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ExternalID, "12ab34")
			},
		},
		"should successfully search for value in array fields": {
			givenKey:         "tags",
			givenSearchValue: "Marshall Islands",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
			},
		},
		"should successfully search for value with spaces": {
			givenKey:         "subject",
			givenSearchValue: "A Problem in South Africa",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Subject, "A Problem in South Africa")
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
		"should return ticket for case insensitive searches - all lower": {
			givenKey:         "priority",
			givenSearchValue: "high",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Priority, "high")
			},
		},
		"should return ticket for case insensitive searches - all upper": {
			givenKey:         "priority",
			givenSearchValue: "HIGH",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Priority, "high")
			},
		},
		"should return ticket for case insensitive searches -mixed": {
			givenKey:         "priority",
			givenSearchValue: "HiGh",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Priority, "high")
			},
		},
		"should return ticket for case insensitive searches in array fields - all lower": {
			givenKey:         "tags",
			givenSearchValue: "georgia",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Tags[0], "Georgia")
			},
		},
		"should return assigned user and related organisation for the ticket": {
			givenKey:         "ID",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(response *model.Response, err error) {
				tkts := response.Tickets
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
				require.Equal(t, tkts[0].AssigneeID, 1)
				require.Equal(t, tkts[0].OrganizationID, 101)

				users := response.Users
				require.Len(t, users, 1)
				require.Equal(t, users[0].Id, 1)

				orgs := response.Orgs
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ID, 101)
			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			tktsearch := NewTicketSearch("testdata/ticket_test.json", "testdata/user_test.json", "testdata/org_test.json")
			tktsearch.Initialise()
			result, err := tktsearch.searchTicket(test.givenKey, test.givenSearchValue)
			test.thenAssert(result, err)
		})
	}
}

func Test_TicketCommands(t *testing.T) {
	tests := map[string]struct {
		thenAssert        func(result int)
		givenArgs         []string
		givenUserFileName string
		givenOrgFileName  string
		givenTktFileName  string
	}{
		"should return success for valid key value": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenTktFileName:  "testdata/ticket_test.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, 0)

			},
		},
		"should return help code when no arguments are passed": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenTktFileName:  "testdata/ticket_test.json",
			givenArgs:         []string{},
			thenAssert: func(result int) {
				require.Equal(t, result, -18511)

			},
		},
		"should return help code when only 1 argument is passed": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenTktFileName:  "testdata/ticket_test.json",
			givenArgs:         []string{},
			thenAssert: func(result int) {
				require.Equal(t, result, -18511)

			},
		},
		"should return success and ignore extra arguments": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenTktFileName:  "testdata/ticket_test.json",
			givenArgs:         []string{"id", "1", "abc"},
			thenAssert: func(result int) {
				require.Equal(t, result, 0)

			},
		},
		"should return -1 for invalid user file": {
			givenUserFileName: "testdata/user.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenTktFileName:  "testdata/ticket_test.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, -1)

			},
		},
		"should return -1 for invalid org file": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org.json",
			givenTktFileName:  "testdata/ticket_test.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, -1)

			},
		},
		"should return -1 for invalid ticket file": {
			givenUserFileName: "testdata/user_test.json",
			givenOrgFileName:  "testdata/org_test.json",
			givenTktFileName:  "testdata/ticket.json",
			givenArgs:         []string{"id", "1"},
			thenAssert: func(result int) {
				require.Equal(t, result, -1)

			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			tktsearch := NewTicketSearch(test.givenTktFileName, test.givenUserFileName, test.givenOrgFileName)
			result := tktsearch.Run(test.givenArgs)
			test.thenAssert(result)
		})
	}
}
