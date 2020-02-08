package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ConvertResponseBodyToStruct(closer io.ReadCloser, obj interface{}) error {
	bytes, err := ioutil.ReadAll(closer)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, obj)
}
