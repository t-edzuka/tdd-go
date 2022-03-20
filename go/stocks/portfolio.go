package stocks

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) Money {
	var totalMoney float64
	// NOT implemented different currency!!
	for _, money := range p {
		totalMoney += money.amount
	}
	return Money{amount: totalMoney, currency: currency}
}
