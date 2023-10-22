package softlinePayment

import (
	"io"
	"time"
)

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

type Signature struct {
	SecretKey     string
	Event         string
	OrderID       string
	CreateDate    string
	PaymentMethod string
	Currency      string
	CustomerEmail string
}

type PaymentResp struct {
	Signature      string    `json:"-"`
	RespBody       []byte    `json:"-"`
	Event          string    `json:"event"`
	EventDate      time.Time `json:"event_date"`
	OrderId        int       `json:"order_id"`
	OrderName      string    `json:"order_name"`
	Status         string    `json:"status"`
	ExternalId     string    `json:"external_id"`
	CreateDate     time.Time `json:"create_date"`
	PayDate        string    `json:"pay_date"`
	Currency       string    `json:"currency"`
	Locale         string    `json:"locale"`
	OrderDetailUrl string    `json:"order_detail_url"`
	Customer       struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
	} `json:"customer"`
	Payment struct {
		Method               string `json:"payment_method"`
		SystemName           string `json:"payment_system_name"`
		ErrorDescription     string `json:"payment_error_description"`
		ErrorCode            string `json:"payment_error_code"`
		CardLast4            int    `json:"card_last_4"`
		CardExpirationDate   string `json:"card_expiration_date"`
		IsCardExpired        bool   `json:"is_card_expired"`
		IsInstallmentPayment bool   `json:"is_installment_payment"`
	} `json:"payment"`
	Return struct {
		Type   string    `json:"type"`
		Reason string    `json:"reason"`
		Date   time.Time `json:"date"`
	} `json:"return"`
}
