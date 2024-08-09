package jsonrpchttp

import (
	"net/http"
)

type Client struct {
	Http Httper
}

type Httper interface {
	Get(url string, options ...OptFunc) (Response, error)
}

type OptFunc func(*Request) (*Request, error)
type Request struct{ *http.Request }
type Response struct{ *http.Response }

func NewClient() (c *Client, err error) {
	return
}
