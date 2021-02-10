package pointers

import "testing"

func TestCarteira(t *testing.T) {
	verifyOutput := func(t *testing.T, carteira Carteira, waited Bitcoin) {
		t.Helper()
		result := carteira.Saldo()
		if result != waited {
			t.Errorf("result '%s', waited '%s'", result, waited)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		waited := Bitcoin(10)

		verifyOutput(t, carteira, waited)
	})

	t.Run("Cashout", func(t *testing.T) {
		carteira := Carteira{money: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		waited := Bitcoin(10)

		verifyOutput(t, carteira, waited)
	})

	t.Run("withdraw with insufficient balance", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		verifyOutput(t, carteira, saldoInicial)
		if erro == nil {
			t.Errorf("Esperava um erro")
		}
	})
}
