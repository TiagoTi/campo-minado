package main

import (
	"campo-minado/aplicacao/dominio"
	"campo-minado/aplicacao/porta/saida"
	jogoAplic "campo-minado/aplicacao/servicos/jogo"
	mock_saida "campo-minado/mocks/aplicacao/porta/saida"
	"fmt"
	"math/rand"
)

var sorteador dominio.GeradorPosicoes
var repositorio saida.RepositorioCampoMinado

func main() {
	sorteador = dominio.NovoGeradorPosicoes()
	repositorio = mock_saida.NewMockRepositorioCampoMinado(nil)
	sessao, err := jogoAplic.NovaSessaoCampoMinado(sorteador, repositorio)
	if err != nil {
		panic(err)
	}

	nomePartida := "poc"
	tamanhoTabuleiro := dominio.Posicao{X: rand.Intn(5) + 2, Y: rand.Intn(5) + 3}
	quantidadeMinas := (tamanhoTabuleiro.X * tamanhoTabuleiro.Y) / 5
	jogo, err := sessao.Novo(nomePartida, tamanhoTabuleiro, quantidadeMinas)
	if err != nil {
		panic(err)
	}

	// name, tamanho, qtBomb,
	posicoesIniciais := dominio.Posicao{X: rand.Intn(tamanhoTabuleiro.X), Y: rand.Intn(tamanhoTabuleiro.Y)}
	t, err := sorteador.Tabuleiro(tamanhoTabuleiro, quantidadeMinas, posicoesIniciais)

	// validação
	if err == nil {
		t.Debug()
		fmt.Println()
		t.DebugRaw()
		fmt.Println()
		t.DebugDefualt()
	}

	fmt.Printf("\nEstado do jogo: %+v\n", jogo.Estado)
	// jogo, _ = sessao.Revelar(jogo.ID, dominio.Posicao{X: 2, Y: 1})
}
