package stocks

import (
	"fmt"
)

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func keyConvert(from string, to string) string {
	return fmt.Sprintf("%s->%s", from, to)
}

func convert(money Money, currency string) float64 {

	exchangeRate := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	if money.currency == currency {
		return money.amount
	}
	key := keyConvert(money.currency, currency)
	return money.amount * exchangeRate[key]
}

func (p Portfolio) Evaluate(currency string) Money {
	var totalAmount float64
	// NOT implemented different currency!!
	for _, money := range p {
		totalAmount += convert(money, currency)
	}
	return NewMoney(totalAmount, currency)
}
