package softlinePayment

import "io"

type SendParams struct {
	Path        string
	HttpMethod  string
	Date        string
	Token       string
	AuthNeed    bool
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

type CreatePaymentReq struct {
	Currency           string   `json:"currency"`
	Amount             string   `json:"amount"`
	ReturnSuccessUrl   string   `json:"return_success_url"`
	PaymentMethod      string   `json:"payment_method"`
	RecurringIndicator bool     `json:"recurring_indicator"`
	PaymentId          string   `json:"payment_id"`
	Customer           Customer `json:"customer"`
}

type Customer struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreatePaymentResp struct {
	PaymentUrl string `json:"payment_url"`
	OrderId    int    `json:"order_id"`
}

type MakePaymentReq struct {
	ParentOrderId      int    `json:"parent_order_id"`
	PaymentId          string `json:"payment_id"`
	Currency           string `json:"currency"`
	Amount             string `json:"amount"`
	PaymentDescription string `json:"payment_description"`
}
