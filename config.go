package softlinePayment

type Config struct {
	IdleConnTimeoutSec int
	RequestTimeoutSec  int
	Login              string
	Pass               string
	URI                string
}
