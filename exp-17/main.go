package main

type Conta struct {
	saldo int
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) Depositar(valor int) {
	c.saldo += valor
}

func main() {
	conta := newConta()
	conta.Depositar(100)
	println(conta.saldo)
}
