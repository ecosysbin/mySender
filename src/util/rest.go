package util

import (
	"fmt"
	"net/http"
)

var client *http.Client

func NewClient() (error, *http.Client) {
	fmt.Println("get client")
	return nil, http.DefaultClient
}

func NewRequest() {

}

func SendRequest() {

}
