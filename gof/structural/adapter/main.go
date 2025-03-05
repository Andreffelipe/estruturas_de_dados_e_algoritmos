package main

import "fmt"

// Interface para processar pagamentos
type ProcessadorPagamento interface {
	realizarPagamento(valor float64, conta string) bool
	verificarFundos(conta string) bool
}

// Implementação do sistema legado
type SistemaLegado struct{}

func (s *SistemaLegado) efetuarTransacao(valor float64, conta string) bool {
	return true
}

func (s *SistemaLegado) consultarSaldoDisponivel(conta string) float64 {
	return 1000.00
}

// Adaptador para o sistema legado para o ProcessadorPagamento
type AdaptadorPagamento struct {
	sistemaLegado *SistemaLegado
}

func (a *AdaptadorPagamento) realizarPagamento(valor float64, conta string) bool {
	return a.sistemaLegado.efetuarTransacao(valor, conta)
}

func (a *AdaptadorPagamento) verificarFundos(conta string) bool {
	return a.sistemaLegado.consultarSaldoDisponivel(conta) > 0
}

// Serviço de cobrança
type ServicoCobranca struct {
	processador ProcessadorPagamento
}

func (s *ServicoCobranca) cobrar(valor float64, conta string) bool {
	if s.processador.verificarFundos(conta) {
		sucesso := s.processador.realizarPagamento(valor, conta)
		if sucesso {
			fmt.Println("Pagamento realizado com sucesso.")
			return true
		} else {
			fmt.Println("Pagamento falhou.")
		}
	}
	return false
}

func main() {
	processador := &AdaptadorPagamento{&SistemaLegado{}}
	servico := ServicoCobranca{processador}
	servico.cobrar(100.0, "123456")
	servico.cobrar(1001.0, "123455")
}
