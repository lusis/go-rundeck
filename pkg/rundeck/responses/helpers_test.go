package responses

import (
	"io/ioutil"
	"os"
)

func testReadJSON(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	return ioutil.ReadAll(file)
}
