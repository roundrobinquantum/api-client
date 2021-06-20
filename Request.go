package api_client

import (
	"encoding/json"
	"fmt"
	"github.com/roundrobinquantum/api-client/errors"
	"github.com/valyala/fasthttp"
)

const (
	CONTENT_TYPE_NAME = "contentType" //todo changeble name
	APP_JSON_NAME     = "application/json"
)

type headerConf struct {
	header string
	value  string
}

type Request struct {
	method  string
	url     string
	body    interface{}
	headers []headerConf
}

func Get(url string) *Request {
	return &Request{
		method: fasthttp.MethodGet,
		url:    url,
	}
}

func Delete(url string) *Request {
	return &Request{
		method: fasthttp.MethodDelete,
		url:    url,
	}
}

func (r *Request) WithHeader(header string, value interface{}) *Request {
	if value == nil {
		return r
	}

	r.headers = append(r.headers, headerConf{header: header, value: value.(string)})
	r.headers = append(r.headers, headerConf{header: CONTENT_TYPE_NAME, value: APP_JSON_NAME})
	return r
}

func (r *Request) WithQuery(k, v string) *Request {
	r.url = fmt.Sprintf("%s%s=%s&", r.url, k, v)
	return r
}

func (r *Request) Build() Request {
	return *r
}

func(r *Request) convertToFastHttpRequest() (*fasthttp.Request, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(r.url)
	req.Header.SetMethod(r.method)

	for _, headerValue := range r.headers {
		req.Header.Set(headerValue.header, headerValue.value)
	}

	if r.body != nil {
		body, err := json.Marshal(r.body)
		if err != nil {
			error := errors.DefineError("API-ERROR",500, fmt.Sprintf(" an error occured when requesting the url : is %s, error definition is : %s", r.url, err.Error()))
			errors.Panic(error)
		}
		req.SetBody(body)
	}

	return req, nil
}


