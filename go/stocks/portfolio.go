package stocks

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func convert(money Money, currency string) float64 {
	euroToUSD := 1.2
	if money.currency == currency {
		return money.amount
	}
	return money.amount * euroToUSD
}

func (p Portfolio) Evaluate(currency string) Money {
	var totalAmount float64
	// NOT implemented different currency!!
	for _, money := range p {
		totalAmount += convert(money, currency)
	}
	return NewMoney(totalAmount, currency)
}
