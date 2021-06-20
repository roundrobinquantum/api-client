## *API CLIENT*

This library supports only get, delete and post with multiple headers

The client coded with fluent language

##### *with header*
```
func (r *Request) WithHeader(header string, value interface{}) *Request {
	if value == nil {
		return r
	}

	r.headers = append(r.headers, headerConf{header: header, value: value.(string)})
	r.headers = append(r.headers, headerConf{header: CONTENT_TYPE_NAME, value: APP_JSON_NAME})
	return r
}

with version 0.0.1

```