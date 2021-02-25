package reader

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_FileReader(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	tests := map[string]struct {
		thenAssert                func(b []byte, err error)
		givenFileName func() string
	}{
		"should successfully read a file passed": {
			givenFileName: func() string {
				return "testdata/sample.json"
			},
			thenAssert: func(b []byte, err error) {
				require.Nil(t, err)
				require.NotNil(t, b)
			},
		},
		"should  error for a file that doesn't exist": {
			givenFileName: func() string {
				return "testdata/nonexistantfile.json"
			},
			thenAssert: func(b []byte, err error) {
				require.NotNil(t, err)
			},
		},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {

			reader:= NewReader()
			b,err := reader.ReadFixtureFile(test.givenFileName())
			test.thenAssert(b, err)
		})
	}
}
