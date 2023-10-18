package decorator

const (
	PaymentSuccess = "success"
	PaymentFail    = "fail"

	commission float64 = 5
)

type PaymentSystem interface {
	Pay(float64) string
}

type DefaultPayment struct {
	amount float64
}

func (p *DefaultPayment) Pay(sum float64) string {
	if p.amount < sum {
		return PaymentFail
	}
	p.amount -= sum
	return PaymentSuccess
}

func NewPSWithCommission(amount float64) PaymentSystem {
	return &PaymentWithCommission{ps: &DefaultPayment{amount: amount}}
}

type PaymentWithCommission struct {
	ps PaymentSystem
}

func (p *PaymentWithCommission) Pay(sum float64) string {
	return p.ps.Pay(sum + commission)
}
