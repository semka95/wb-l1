package main

import (
	"fmt"
)

type Currency struct {
	amount int
}

func (c *Currency) GetUSD() int {
	return c.amount
}

func (c *Currency) SetUSD(amount int) {
	c.amount = amount
}

type Money interface {
	GetRUB() int
	SetRUB(int)
}

type MoneyAdapter struct {
	curr Currency
}

func (m *MoneyAdapter) GetRUB() int {
	return m.curr.GetUSD() * 300
}

func (m *MoneyAdapter) SetRUB(amount int) {
	m.curr.SetUSD(amount / 300)
}

func main() {
	var rub Money = &MoneyAdapter{
		curr: Currency{100},
	}
	fmt.Printf("initial amount of money:%d rub\n", rub.GetRUB())
	rub.SetRUB(300)
	fmt.Printf("money:%d rub\n", rub.GetRUB())
}
