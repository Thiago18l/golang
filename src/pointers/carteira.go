package pointers

// Carteira struct
type Carteira struct {
	money int
}

// Saldo method
func (c *Carteira) Saldo() int {
	return c.money
}

// Depositar method
func (c *Carteira) Depositar(money int) {
	c.money += money
}
