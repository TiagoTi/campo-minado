package gamesrv_test

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func RN001(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)

	var (
		newGameError error
		game         *domain.Game
	)

	gameService := gamesrv.NewGameService(mockGamRep)

	assert.NotNil(t, gameService)

	game, newGameError = gameService.Create(0, 0, 0)

	assert.Nil(t, game)
	require.Error(t, newGameError)
	assert.Equal(t, "o tabuleiro do jogo deve ter um quantidade de linhas maior que um e recebeu 0", newGameError.Error())
}
