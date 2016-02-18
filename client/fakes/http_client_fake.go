package client_fakes

import (
	"io"
	"net/http"
	"net/url"
)

type FakeHttpClient struct {
	DoRequest *http.Request

	GetUrl string

	HeadUrl string

	PostUrl      string
	PostBodyType string
	PostBody     io.Reader

	PostFormUrl string
	PostData    url.Values

	FakeResponse *http.Response
	FakeErr      error
}

func NewFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{
		FakeResponse: &http.Response{},
		FakeErr:      nil,
	}
}

func (c *FakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.DoRequest = req

	return c.FakeResponse, c.FakeErr
}

func (c *FakeHttpClient) Get(url string) (*http.Response, error) {
	c.GetUrl = url

	return c.FakeResponse, c.FakeErr
}

func (c *FakeHttpClient) Head(url string) (*http.Response, error) {
	c.HeadUrl = url

	return c.FakeResponse, c.FakeErr
}

func (c *FakeHttpClient) Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	c.PostUrl = url
	c.PostBodyType = bodyType
	c.PostBody = body

	return c.FakeResponse, c.FakeErr
}

func (c *FakeHttpClient) PostForm(url string, data url.Values) (*http.Response, error) {
	c.PostFormUrl = url
	c.PostData = data

	return c.FakeResponse, c.FakeErr
}
