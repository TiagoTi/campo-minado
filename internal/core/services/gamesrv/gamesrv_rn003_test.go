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

func RN003(t *testing.T) {
	var (
		newGameError error
		game         *domain.Game
		lines        int = 2
		columns      int = 2
	)

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
	mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)

	gameService := gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

	assert.NotNil(t, gameService)

	game, newGameError = gameService.Create(lines, columns, 0)

	assert.Nil(t, game)
	require.Error(t, newGameError)
	assert.Equal(t, "uma sessão de jogo deve possuir no mínimo uma mina e recebeu 0", newGameError.Error())
}
