package softlinePayment

import "io"

type SendParams struct {
	Path        string
	HttpMethod  string
	Date        string
	Body        io.Reader
	QueryParams map[string]string
	Response    interface{}
}

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResp struct {
	Token        string
	RefreshToken string
	Date         string
}
