package entrada_test

import (
	"campo-minado/aplicacao/dominio"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNovoGeradorPosicoes(t *testing.T) {
	var tamanho dominio.Posicao
	var qntBomba int
	var posicaoInicial dominio.Posicao
	t.Run("erro tamanho tabuleiro < 2", func(t *testing.T) {
		tamanho = dominio.Posicao{X: 1, Y: 1}
		qntBomba = 1
		posicaoInicial = dominio.Posicao{X: 0, Y: 0}
		tab, e := dominio.NovoGeradorPosicoes().Tabuleiro(tamanho, qntBomba, posicaoInicial)
		assert.Nil(t, tab)
		assert.NotNil(t, e)
		assert.EqualError(t, e, "erro tamanho tabuleiro < 2")
	})
	t.Run("erro qnt bomba igual tamanho tabuleiro", func(t *testing.T) {
		tamanho = dominio.Posicao{X: 2, Y: 2}
		qntBomba = 4
		posicaoInicial = dominio.Posicao{X: 0, Y: 0}
		tab, e := dominio.NovoGeradorPosicoes().Tabuleiro(tamanho, qntBomba, posicaoInicial)
		assert.Nil(t, tab)
		assert.NotNil(t, e)
		assert.EqualError(t, e, "erro qnt bomba igual tamanho tabuleiro")
	})
	t.Run("erro qnt bomba maior tamanho tabuleiro", func(t *testing.T) {
		tamanho = dominio.Posicao{X: 2, Y: 2}
		qntBomba = 5
		posicaoInicial = dominio.Posicao{X: 0, Y: 0}
		tab, e := dominio.NovoGeradorPosicoes().Tabuleiro(tamanho, qntBomba, posicaoInicial)
		assert.Nil(t, tab)
		assert.NotNil(t, e)
		assert.EqualError(t, e, "erro qnt bomba > tamanho tabuleiro")
	})
	t.Run("erro qnt bomba < 1", func(t *testing.T) {
		tamanho = dominio.Posicao{X: 2, Y: 2}
		qntBomba = 0
		posicaoInicial = dominio.Posicao{X: 0, Y: 0}
		tab, e := dominio.NovoGeradorPosicoes().Tabuleiro(tamanho, qntBomba, posicaoInicial)
		assert.Nil(t, tab)
		assert.NotNil(t, e)
		assert.EqualError(t, e, "erro qntBomba < 1")
	})
	t.Run("celula inicial deve ser igual posicao iniciao", func(t *testing.T) {
		tamanho = dominio.Posicao{X: 2, Y: 2}
		qntBomba = 1
		posicaoInicial = dominio.Posicao{X: 0, Y: 0}
		tab, e := dominio.NovoGeradorPosicoes().Tabuleiro(tamanho, qntBomba, posicaoInicial)
		assert.Equal(t, 2, len(tab))
		assert.Nil(t, e)
		assert.Equal(t, dominio.EstadoCelulaAberta, tab[posicaoInicial.X][posicaoInicial.Y])
	})
}
