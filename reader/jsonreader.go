package reader

import (
	"io/ioutil"
	"os"
)

type Reader interface {
	ReadFixtureFile( fileName string) ([]byte, error)
}

func NewReader() Reader {
	return &reader{
	}
}

type reader struct{
}

func (r *reader) ReadFixtureFile( fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

