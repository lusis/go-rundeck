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
	// This converts our custom JSONTime/JSONDuration properly
	jsonTimeHook := func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		// we only work on strings
		if f.Kind() != reflect.String {
			return data, nil
		}

		// first we check if it's a JSONTime
		if t == reflect.TypeOf(&JSONTime{time.Time{}}) {
			// Convert it by parsing
			tTime, err := time.Parse(rdTime, data.(string))
			if err != nil {
				return nil, err
			}
			return JSONTime{tTime}, nil
		}

		// Next we check if it's a JSONDuration
		sampleDuration, _ := time.ParseDuration("1s")
		if t == reflect.TypeOf(&JSONDuration{sampleDuration}) {
			dur, err := time.ParseDuration(data.(string))
			if err != nil { return nil, err}
			return JSONDuration{dur}, nil
		}

		return data, nil
	}
	return &mapstructure.DecoderConfig{
		ErrorUnused:      true,
		WeaklyTypedInput: false,
		TagName:          "json",
		DecodeHook:       jsonTimeHook,
	}
}
