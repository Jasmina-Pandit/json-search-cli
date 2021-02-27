package cmd

import (
	"github.com/stretchr/testify/require"
	"json-search-cli/model"
	"testing"
)

func Test_TicketSearch(t *testing.T) {

	tests := map[string]struct {
		thenAssert       func(tkts []model.Ticket, err error)
		givenKey         string
		givenSearchValue string
	}{
		"should successfully search small case id": {
			givenKey:         "id",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search upper case id": {
			givenKey:         "ID",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search underscore id": {
			givenKey:         "_id",
			givenSearchValue: "50f3fdbd-f8a6-481d-9bf7-572972856628",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search by external_id": {
			givenKey:         "external_id",
			givenSearchValue: "12ab34",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ID, "50f3fdbd-f8a6-481d-9bf7-572972856628")
			},
		},
		"should successfully search by externalid without underscore": {
			givenKey:         "externalid",
			givenSearchValue: "12ab34",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].ExternalID, "12ab34")
			},
		},
		"should successfully search for value in array fields": {
			givenKey:         "tags",
			givenSearchValue: "Marshall Islands",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
			},
		},
		"should successfully search for value with spaces": {
			givenKey:         "subject",
			givenSearchValue: "A Problem in South Africa",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.NotNil(t, tkts)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Subject, "A Problem in South Africa")
			},
		},
		"should not error on missing fields": {
			givenKey:         "newkey",
			givenSearchValue: "hello",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.NotNil(t, err)
				require.Error(t, err, "invalid key. Use help command for list of valid keys")

			},
		},
		"should return user for case insensitive searches - all lower": {
			givenKey:         "priority",
			givenSearchValue: "high",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Priority, "high")
			},
		},
		"should return user for case insensitive searches - all upper": {
			givenKey:         "priority",
			givenSearchValue: "HIGH",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Priority, "high")
			},
		},
		"should return user for case insensitive searches -mixed": {
			givenKey:         "priority",
			givenSearchValue: "HiGh",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Priority, "high")
			},
		},
		"should return user for case insensitive searches in array fields - all lower": {
			givenKey:         "tags",
			givenSearchValue: "georgia",
			thenAssert: func(tkts []model.Ticket, err error) {
				require.Nil(t, err)
				require.Len(t, tkts, 1)
				require.Equal(t, tkts[0].Tags[0], "Georgia")
			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			tktsearch := NewTicketSearch("testdata/ticket_test.json")
			result, err := tktsearch.searchTicket(test.givenKey, test.givenSearchValue)
			test.thenAssert(result, err)
		})
	}
}
