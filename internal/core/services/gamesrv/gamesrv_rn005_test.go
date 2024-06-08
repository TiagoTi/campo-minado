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

func RN005(t *testing.T) {
	var (
		newGameError error
		game         *domain.Game
		lines        int = 3
		columns      int = 3
		mines        int = 1
	)

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
	mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)

	gameService := gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

	assert.NotNil(t, gameService)

	game, newGameError = gameService.Create(lines, columns, mines)

	assert.Nil(t, game)
	require.Error(t, newGameError)
	assert.Equal(t, "-", newGameError.Error())
}
