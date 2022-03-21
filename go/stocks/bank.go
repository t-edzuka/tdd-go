package stocks

import "errors"

type Bank struct {
	exchangeRates map[string]float64
}

func NewBank() Bank {
	return Bank{
		exchangeRates: make(map[string]float64),
	}
}

func getConvertKey(currencyFrom string, currencyTo string) string {
	return currencyFrom + "->" + currencyTo
}

func (b *Bank) AddExchangeRate(currencyFrom string, currencyTo string, rate float64) {
	key := getConvertKey(currencyFrom, currencyTo)
	b.exchangeRates[key] = rate
}

func (b *Bank) Convert(money Money, currencyTo string) (convertedMoney *Money,
	err error) {
	if money.currency == currencyTo {
		result := NewMoney(money.amount, money.currency)
		return &result, nil
	}
	key := getConvertKey(money.currency, currencyTo)

	rate, isOk := b.exchangeRates[key]
	if !isOk {
		return nil, errors.New(key)
	}
	result := NewMoney(money.amount*rate, currencyTo)
	return &result, nil
}
