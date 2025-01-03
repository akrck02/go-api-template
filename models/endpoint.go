package models

type Endpoint struct {
	Path   string     `json:"path"`
	Method HttpMethod `json:"method"`

	RequestMimeType  MimeType `json:"requestMimeType"`
	ResponseMimeType MimeType `json:"responseMimeType"`

	Listener EndpointListener `json:"-"`
	Checks   EndpointCheck    `json:"-"`

	Secured  bool `json:"secured"`
	Database bool `json:"-"`
}

type EndpointCheck func(context *ApiContext) *Error
type EndpointListener func(context *ApiContext) (*interface{}, *Error)
