package servicos_test

import (
	"campo-minado/aplicacao/dominio"
	"campo-minado/aplicacao/porta/entrada"
	jogoAplic "campo-minado/aplicacao/servicos/jogo"
	mock_saida "campo-minado/mocks/aplicacao/porta/saida"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	sessaoCampoMinado entrada.SessaoCampoMinado
)

const (
	m = dominio.EstadoCelulaMina
	o = dominio.EstadoCelulaOculta
	a = dominio.EstadoCelulaAberta
	b = dominio.EstadoCelulaBandeira
)

func Test_Revel(t *testing.T) {
	t.Run("o usuario não pode escolher uma linha maior que a quantidade de linha do tabuleiro", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		jogoMock := &dominio.CampoMinado{Estado: dominio.EstadoJogoAtivo, Tamanho: dominio.Posicao{X: 2, Y: 2}}
		mockRepo.EXPECT().Obter("01").Return(jogoMock, nil)
		s, _ := jogoAplic.NovaSessaoCampoMinado(nil, mockRepo)
		_, e := s.Revelar("01", dominio.Posicao{X: 0, Y: 3})
		assert.EqualError(t, e, "a linha deve estar entre 0 e 2")
	})
	t.Run("deve dar erro quando o usuario escolher uma linha menor que 0", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		jogoMock := &dominio.CampoMinado{Estado: dominio.EstadoJogoAtivo}
		mockRepo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		mockRepo.EXPECT().Obter("01").Return(jogoMock, nil)
		s, _ := jogoAplic.NovaSessaoCampoMinado(nil, mockRepo)
		_, e := s.Revelar("01", dominio.Posicao{X: 0, Y: -1})
		assert.EqualError(t, e, "a linha deve iniciar da posicao 0")
	})

	t.Run("o usuario não pode escolher uma coluna maior que a quantidade de colunas do tabuleiro", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		jogoMock := &dominio.CampoMinado{Estado: dominio.EstadoJogoAtivo, Tamanho: dominio.Posicao{X: 2, Y: 2}}
		mockRepo.EXPECT().Obter("01").Return(jogoMock, nil)
		s, _ := jogoAplic.NovaSessaoCampoMinado(nil, mockRepo)
		_, e := s.Revelar("01", dominio.Posicao{X: 3, Y: 0})
		assert.EqualError(t, e, "a coluna deve estar entre 0 e 2")
	})
	t.Run("deve dar erro quando o usuario escolher uma coluna menor que 0", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		jogoMock := &dominio.CampoMinado{Estado: dominio.EstadoJogoAtivo}
		mockRepo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		mockRepo.EXPECT().Obter("01").Return(jogoMock, nil)
		s, _ := jogoAplic.NovaSessaoCampoMinado(nil, mockRepo)
		_, e := s.Revelar("01", dominio.Posicao{X: -1, Y: 0})
		assert.EqualError(t, e, "a coluna deve iniciar da posicao 0")
	})
	t.Run("retorna erro ao tentar revelar um jogo com estado de ganho", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repositorioJogo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		tabuleiro := &dominio.Tabuleiro{{m}}
		mock := &dominio.CampoMinado{Estado: dominio.EstadoJogoGanho, Tabuleiro: *tabuleiro}
		repositorioJogo.EXPECT().Obter("01").Return(mock, nil)
		sessaoCampoMinado, _ = jogoAplic.NovaSessaoCampoMinado(nil, repositorioJogo)
		_, e := sessaoCampoMinado.Revelar("01", dominio.Posicao{X: 0, Y: 0})
		assert.EqualError(t, e, "este jogo já foi encerrado")
	})
	t.Run("retorna erro ao tentar revelar um jogo com estado de perda", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repositorioJogo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		tabuleiro := &dominio.Tabuleiro{{m}}
		mock := &dominio.CampoMinado{Estado: dominio.EstadoJogoPerda, Tabuleiro: *tabuleiro}
		repositorioJogo.EXPECT().Obter("01").Return(mock, nil)
		sessaoCampoMinado, _ = jogoAplic.NovaSessaoCampoMinado(nil, repositorioJogo)
		_, e := sessaoCampoMinado.Revelar("01", dominio.Posicao{X: 0, Y: 0})
		assert.EqualError(t, e, "este jogo já foi encerrado")
	})
	t.Run("o estado do jogo deve ser perda quando revela uma bomba", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repositorioJogo := mock_saida.NewMockRepositorioCampoMinado(ctrl)
		tabuleiro := &dominio.Tabuleiro{{m}}
		mock := &dominio.CampoMinado{Estado: dominio.EstadoJogoAtivo, Tabuleiro: *tabuleiro}
		repositorioJogo.EXPECT().Obter("01").Return(mock, nil)
		sessaoCampoMinado, _ = jogoAplic.NovaSessaoCampoMinado(nil, repositorioJogo)
		j, _ := sessaoCampoMinado.Revelar("01", dominio.Posicao{X: 0, Y: 0})
		assert.Equal(t, j.Estado, dominio.EstadoJogoPerda)
	})
}

func Fuzz_NovaSessaoCampoMinado(f *testing.F) {
	testcases := []struct {
		nome                          string
		qntColuna, qntLinha, qntBomba int
		mensgErr                      string
	}{
		{"", -1, -1, -1, "nome não informado"},
		{"jogo", -1, -1, -1, "o jogo não pode ter quantidade negativa de coluna"},
		{"jogo", 0, -1, -1, "o jogo dever ter uma quantidade de coluna maior que zero"},
		{"jogo", 1, -1, -1, "o jogo dever ter uma quantidade de coluna maior que 1"},
		{"jogo", 2, -1, -1, "o jogo não pode ter quantidade negativa de linha"},
		{"jogo", 2, 0, -1, "o jogo dever ter uma quantidade de linha maior que zero"},
		{"jogo", 2, 1, -1, "o jogo dever ter uma quantidade de linha maior que 1"},
		{"jogo", 2, 2, -1, "quantidade de bombas não pode ser negativa"},
		{"jogo", 2, 2, 0, "o jogo deve ter pelo menos uma bomba"},
		{"jogo", 2, 2, 5, "a quantidade de bombas deve ser menor que o total de espaço do tabuleiro"},
		{"jogo", 2, 2, 4, "a quantidade de bombas deve ser menor que o total de espaço do tabuleiro"},
		{"jogo", 2, 2, 1, ""},
	}
	for _, tc := range testcases {
		f.Add(tc.nome, tc.qntColuna, tc.qntLinha, tc.qntBomba, tc.mensgErr)
	}
	f.Fuzz(func(t *testing.T, nome string, qntColuna, qntLinha, qntBomba int, erro string) {
		s, e := jogoAplic.NovaSessaoCampoMinado(nil, nil)
		assert.NotNil(t, s)
		assert.Nil(t, e)
		tamanhoTabuleiro := dominio.Posicao{X: qntColuna, Y: qntLinha}
		j, e := s.Novo(nome, tamanhoTabuleiro, qntBomba)
		if e != nil {
			assert.Nil(t, j)
			assert.Equal(t, erro, e.Error())
		} else {
			assert.Equal(t, nome, j.Nome)
			assert.NotNil(t, j)
			assert.Equal(t, "", erro)
		}
	})
}
