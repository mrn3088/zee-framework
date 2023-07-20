package web_framework

import "net/http"

// Context context
type Context struct {
	// response
	response http.ResponseWriter
	// request
	request *http.Request
	// Method request method
	Method string
	// Request URL
	Pattern string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,
		Method:   r.Method,
		Pattern:  r.URL.Path,
	}
}
