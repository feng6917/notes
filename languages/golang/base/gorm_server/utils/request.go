package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func UnmarshalBody(ir io.Reader, req interface{})  error{
	var b []byte
	var err error
	b, err = ioutil.ReadAll(ir)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalMap(ir io.Reader,m interface{}) error {
	var b []byte
	var err error
	b, err = ioutil.ReadAll(ir)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &m)
	return err
}
