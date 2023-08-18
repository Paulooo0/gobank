package util

const (
	USD = "USD"
	BRL = "BRL"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, BRL, EUR:
		return true
	}
	return false
}
