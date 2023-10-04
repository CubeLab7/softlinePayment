package softlinePayment

import "io"

type SendParams struct {
	Path        string
	HttpMethod  string
	Body        io.Reader
	QueryParams map[string]string
	Response    interface{}
}

type AuthResp struct {
	Token string
}
