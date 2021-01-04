package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.cash > amount {
		w.cash -= amount
		fmt.Println("Payment success!")
		return nil
	}
	return fmt.Errorf("cash in not enough")
}

type Card struct {
	Balance                             int
	ValidUntil, Cardholder, CVV, Number string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает денег на карте")
	}
	c.Balance -= amount
	return nil
}

func Buy(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не не является платежным средством\n\n", in)
		return
	}
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet := &Wallet{cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "rvasily"}
	Buy(myMoney)

	// myMoney = &ApplePay{Money: 9}
	// Buy(myMoney)
}

/*
guru
vet
godef
dlv
*/
