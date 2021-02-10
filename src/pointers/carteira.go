package pointers

import (
	"errors"
	"fmt"
)

var errorFunds = errors.New("não é possível retirar: saldo insuficiente")

// Stringer interface
type Stringer interface {
	String() string
}

// Bitcoin type
type Bitcoin int

// Carteira struct
type Carteira struct {
	money Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Saldo method
func (c *Carteira) Saldo() Bitcoin {
	return c.money
}

// Depositar method
func (c *Carteira) Depositar(money Bitcoin) {
	c.money += money
}

// Retirar method
func (c *Carteira) Retirar(money Bitcoin) error {
	if money > c.Saldo() {
		return errorFunds
	}
	c.money -= money
	return nil
}
