package dominio

import (
	"fmt"
	"math/rand"
)

type GeradorPosicoes interface {
	Tabuleiro(tamanho Posicao, qntBomba int, posicaoInicial Posicao) (Tabuleiro, error)
}

// NovoGeradorPosicoes novo_gerador_posicoes_test.go
func NovoGeradorPosicoes() GeradorPosicoes {
	return &gerador{}
}

type gerador struct{}

// Tabuleiro define um tabuleiro ignorando a posição informada
func (g *gerador) Tabuleiro(tamanho Posicao, qntBomba int, posicaoInicial Posicao) (Tabuleiro, error) {
	qntCampos := tamanho.X * tamanho.Y
	if qntBomba < 1 {
		return nil, fmt.Errorf("erro qntBomba < 1")
	}
	if tamanho.X < 2 || tamanho.Y < 2 {
		return nil, fmt.Errorf("erro tamanho tabuleiro < 2")
	}
	if qntBomba == qntCampos {
		return nil, fmt.Errorf("erro qnt bomba igual tamanho tabuleiro")
	}
	if qntBomba > qntCampos {
		return nil, fmt.Errorf("erro qnt bomba > tamanho tabuleiro")
	}
	tabuleiro := make(Tabuleiro, tamanho.X)
	for i := range tabuleiro {
		tabuleiro[i] = make([]EstadoCelula, tamanho.Y)
	}

	tabuleiro[posicaoInicial.X][posicaoInicial.Y] = EstadoCelulaAberta
	posicoesAleatorias := embaralhar(qntCampos, qntBomba)
	var linha, coluna, qntPreenchida int
	for _, posicao := range posicoesAleatorias {
		if qntPreenchida == qntBomba {
			// fmt.Println("qntPreenchida == qntBomba")
			break
		}
		linha = posicao / tamanho.Y
		coluna = posicao - linha*tamanho.Y
		// fmt.Printf("l:%d p:%d\n", linha, posicao)
		if linha == posicaoInicial.X && coluna == posicaoInicial.Y {
			// fmt.Println("linha == posicaoInicial.X && coluna == posicaoInicial.Y")
			continue
		}

		qntPreenchida += 1
		tabuleiro[linha][coluna] = EstadoCelulaMina
	}
	// tabuleiro.Debug()
	return tabuleiro, nil
}

func embaralhar(qntCampo, qntBomba int) []int {
	qntBomba += 1
	permutacaoPos := rand.Perm(qntCampo)
	posicoesAleatorias := make([]int, qntBomba)
	// should use copy(to, from) instead of a loop (S1001)go-staticcheck
	// for i, posicao := range permutacaoPos[:qntBomba] {
	// 	posicoesAleatorias[i] = posicao
	// }
	copy(posicoesAleatorias, permutacaoPos[:qntBomba])
	return posicoesAleatorias
}
