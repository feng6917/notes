package main

import (
	"fmt"
	"testing"
)

func TestSendGetRequest(t *testing.T) {
	s, err := SendGetRequest("wangcheng")
	fmt.Println(s, err)
}

func TestSendPostBodyRequest(t *testing.T) {
	s, err := SendPostBodyRequest("wangchen post-body")
	fmt.Println(s, err)
}

func TestSendPostFormRequest(t *testing.T) {
	s, err := SendPostFormRequest("wangchen post-form")
	fmt.Println(s, err)
}

func TestSendPostFormFileRequest(t *testing.T) {
	fpath := "./main.go"
	s, err := SendPostFormFileRequest(fpath)
	fmt.Println(s, err)
}
