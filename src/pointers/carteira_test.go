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

	verifyError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("Esperava um erro")
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
		verifyError(t, erro)
	})
}

func confirmError(t *testing.T, result error, waited error) {
	t.Helper()
	if result == nil {
		t.Fatal("Esperava um erro, mas nenhum aconteceu")
	}
	if result != waited {
		t.Errorf("error result '%s', error waited '%s'", result, waited)
	}

}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}
