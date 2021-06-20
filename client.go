package api_client

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Client struct {
	baseClient *fasthttp.Client
	timeOut    time.Duration
}

func NewClient(timeOut, readTimeout, writeTimeout int) *Client {
	return &Client{
		timeOut: time.Duration(timeOut) * time.Second,
		baseClient: &fasthttp.Client{
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		},
	}
}

func (c *Client) End(r Request) (int, []byte, error) {
	req, err := r.convertToFastHttpRequest()
	if err != nil {
		return 0, nil, err
	}

	resp := fasthttp.AcquireResponse()

	err = c.baseClient.DoTimeout(req, resp, c.timeOut)

	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}