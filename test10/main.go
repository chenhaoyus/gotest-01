package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	creditCard := CreditCard{balance: 500, limit: 2000}

	purchaseItem(&creditCard, 1000)

	fmt.Println("当前余额:", creditCard.GetBalance())

	ditCard := DebitCard{balance: 2000}

	purchaseItem(&ditCard, 1500)

	fmt.Println("当前余额:", ditCard.GetBalance())

}

type PayMethod interface {
	Account
	Pay(amount int) bool
}

type Account interface {
	GetBalance() int
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) GetBalance() int {
	return c.limit - c.balance
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount > c.limit {
		c.balance += amount
		fmt.Println("Payment declined: limit exceeded")
		return false
	} else {
		c.balance += amount
		fmt.Println("Payment successful")
		return true
	}
}

func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

type DebitCard struct {
	balance int
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance < 0 {
		fmt.Println("Payment declined: insufficient funds")
		return false
	} else if d.balance-amount < 0 {
		fmt.Println("Payment declined: insufficient funds")
		return false
	} else {
		d.balance -= amount
		fmt.Println("Payment successful")
		return true
	}
}
