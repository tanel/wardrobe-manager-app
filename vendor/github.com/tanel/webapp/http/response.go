package http

// Response represents response
type Response struct {
	templateName string
	data         interface{}
	redirectURL  string
}

// NewPage returns new page
func NewPage(templateName string, data interface{}) *Response {
	return &Response{
		templateName: templateName,
		data:         data,
	}
}

// NewRedirect returns new redirect
func NewRedirect(redirectURL string) *Response {
	return &Response{
		redirectURL: redirectURL,
	}
}
