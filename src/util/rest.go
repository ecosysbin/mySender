package util

import (
	"net/http"
	"crypto/tls"
	"io"
	"fmt"
)

type Client struct {
	Client *http.Client
}

func NewClient() *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := Client{&http.Client{Transport:tr}}
	return &client
}

func NewRequest(method string, url string, body io.Reader, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return  nil, err
	}
	if headers != nil {
		fmt.Println("header is not nil")
		for headerName, value := range headers {
			request.Header.Set(headerName,value)
		}
	}
	return request, err
}

func (c *Client) SendRequest(r *http.Request) (*http.Response, error) {
	// res := http.Response{}
	res, err := c.Client.Do(r)
	return res, err
}
