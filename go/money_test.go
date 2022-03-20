package main

import (
	"testing"
)

type Multiplier interface {
	Times()
}

type Divider interface {
	divide()
}

type Dollar struct {
	amount int
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}

func TestMultiplication(t *testing.T) {
	fiver := Dollar{
		amount: 5,
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenner.amount)
	}
}

// End of Chapter1

// Start of Chapter2
type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{m.amount * float64(multiplier),
		m.currency}
}

func (m Money) divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %v Got %+v", expected, actual)
	}
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)
	expectedResult := Money{amount: 20, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

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

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	expectedMoney := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, expectedMoney, portfolioInDollars)

}
