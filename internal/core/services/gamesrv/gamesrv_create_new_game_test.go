package gamesrv_test

import (
	"testing"
)

func TestCreateNewGame(t *testing.T) {
	t.Parallel()

	t.Run("criar jogo com sucesso", CreateNewGameSucess)
	t.Run("criar jogo com erro de gravação", CreateNewGameErrorDatabase)

	t.Run("O tabuleiro do jogo deve ter um quantidade de linhas maior que um", RN001)

	t.Run("O tabuleiro do jogo deve ter uma quantidade de colunas maior que um.", RN002)

	t.Run("Uma sessão de jogo deve possuir no mínimo uma mina.", RN003)

	t.Run("uma sessão deve ter menos do que linhas X colunas minas", RN004)
}
