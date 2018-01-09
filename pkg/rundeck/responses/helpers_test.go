package responses

import (
	"io/ioutil"
	"os"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

func testReadJSON(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	return ioutil.ReadAll(file)
}

// newMSDecoderConfig returns a new mapstructure.DecoderConfig we can use for strict validation
// in testing
func newMSDecoderConfig() *mapstructure.DecoderConfig {
	// This converts our custom JSONTime properly
	jsonTimeHook := func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(&JSONTime{time.Now()}) {
			return data, nil
		}

		// Convert it by parsing
		tTime, tErr := time.Parse(rdTime, data.(string))
		if tErr != nil {
			return nil, tErr
		}
		return JSONTime{tTime}, nil
	}
	return &mapstructure.DecoderConfig{
		ErrorUnused:      true,
		WeaklyTypedInput: false,
		TagName:          "json",
		DecodeHook:       jsonTimeHook,
	}
}
