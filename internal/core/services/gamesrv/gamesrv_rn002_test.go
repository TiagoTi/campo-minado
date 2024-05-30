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

func RN002(t *testing.T) {
	var (
		newGameError error
		game         *domain.Game
		lines        uint = 2
	)

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)

	gameService := gamesrv.NewGameService(mockGamRep)

	assert.NotNil(t, gameService)

	game, newGameError = gameService.Create(lines, 0, 0)
	assert.Nil(t, game)
	require.Error(t, newGameError)
	assert.Equal(t, "deve ter pelo menos 2 colunas. Recebeu 0", newGameError.Error())
}
