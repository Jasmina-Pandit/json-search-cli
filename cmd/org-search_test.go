package cmd

import (
	"github.com/stretchr/testify/require"
	"json-search-cli/model"
	"testing"
)

func Test_OrgSearch(t *testing.T) {

	tests := map[string]struct {
		thenAssert       func(response *model.Response, err error)
		givenKey         string
		givenSearchValue string
	}{
		"should successfully search small case id": {
			givenKey:         "id",
			givenSearchValue: "102",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ID, 102)
			},
		},
		"should successfully search upper case id": {
			givenKey:         "ID",
			givenSearchValue: "102",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ID, 102)
			},
		},
		"should successfully search underscore id": {
			givenKey:         "_id",
			givenSearchValue: "102",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ID, 102)
			},
		},
		"should successfully search by external_id": {
			givenKey:         "external_id",
			givenSearchValue: "12ab34",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ID, 102)
			},
		},
		"should successfully search by externalid without underscore": {
			givenKey:         "externalid",
			givenSearchValue: "12ab34",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].ExternalID, "12ab34")
			},
		},
		"should successfully search for value in array fields": {
			givenKey:         "domain_names",
			givenSearchValue: "trollery.com",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
			},
		},
		"should successfully search for value with spaces": {
			givenKey:         "details",
			givenSearchValue: "Non profit",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.NotNil(t, orgs)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].Details, "Non profit")

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
		"should return org for case insensitive searches - all lower": {
			givenKey:         "name",
			givenSearchValue: "nutralab",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].Name, "Nutralab")
			},
		},
		"should return org for case insensitive searches - all upper": {
			givenKey:         "name",
			givenSearchValue: "NUTRALAB",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].Name, "Nutralab")
			},
		},
		"should return org for case insensitive searches - mixed": {
			givenKey:         "name",
			givenSearchValue: "nutrAlaB",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].Name, "Nutralab")
			},
		},
		"should return org for case insensitive searches in array fields - all lower": {
			givenKey:         "tags",
			givenSearchValue: "cherry",
			thenAssert: func(response *model.Response, err error) {
				orgs := response.Orgs
				require.Nil(t, err)
				require.Len(t, orgs, 1)
				require.Equal(t, orgs[0].Tags[0], "Cherry")
			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			orgsearch := NewOrgSearch("testdata/org_test.json")
			orgsearch.Initialise()
			result, err := orgsearch.searchOrg(test.givenKey, test.givenSearchValue)
			test.thenAssert(result, err)
		})
	}
}
