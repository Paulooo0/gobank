package util

const (
	UDS = "USD"
	BRL = "BRL"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case UDS, BRL, EUR:
		return true
	}
	return false
}
