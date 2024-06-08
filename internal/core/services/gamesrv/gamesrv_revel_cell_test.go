package gamesrv_test

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/provides"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRevelCell(t *testing.T) {
	t.Parallel()
	var (
		srv  provides.GamesService
		ctrl *gomock.Controller
	)
	ctrl = gomock.NewController(t)
	ctrl.Finish()
	repo := mocks.NewMockGamesRepository(ctrl)
	generator := mocks.NewMockBoardGenerator(ctrl)
	repoGame := &domain.Game{
		Columns: 3,
		Lines:   3,
		Mines:   1,
	}
	board := domain.Board{
		{0, 0, 0},
		{0, -3, 0},
		{0, 0, 0},
	}
	repo.EXPECT().Get("123").Return(repoGame, nil)
	generator.EXPECT().New(3, 3, 1).Return(board, nil)
	repo.EXPECT().Save(gomock.Any()).Return(repoGame, nil)
	srv = gamesrv.NewGameService(repo, generator)

	game, revealErr := srv.Revel("123", domain.Position{Col: 0, Row: 0})
	game.Debug()

	require.NoError(t, revealErr)
	require.NotNil(t, game)
	assert.Equal(t, game.Status, domain.InProgress)

	oneTip := domain.CellStatus(1)
	hidden := domain.Hidden

	assert.Equal(t, oneTip, game.Board[0][0])
	assert.Equal(t, oneTip, game.Board[0][1])
	assert.Equal(t, oneTip, game.Board[1][0])
	assert.Equal(t, hidden, game.Board[0][2])
	assert.Equal(t, hidden, game.Board[1][2])
	assert.Equal(t, hidden, game.Board[2][0])
	assert.Equal(t, hidden, game.Board[2][1])
	assert.Equal(t, hidden, game.Board[2][2])
	assert.Equal(t, domain.Mine, game.Board[1][1])

}
func TestRevelCell2(t *testing.T) {
	t.Parallel()
	var (
		srv  provides.GamesService
		ctrl *gomock.Controller
	)
	ctrl = gomock.NewController(t)
	ctrl.Finish()
	repo := mocks.NewMockGamesRepository(ctrl)
	generator := mocks.NewMockBoardGenerator(ctrl)
	repoGame := &domain.Game{
		Columns: 3,
		Lines:   3,
		Mines:   1,
	}
	board := domain.Board{
		{1, 1, 0},
		{1, -3, 0},
		{0, 0, 0},
	}
	repo.EXPECT().Get("123").Return(repoGame, nil)
	generator.EXPECT().New(3, 3, 1).Return(board, nil)
	repo.EXPECT().Save(gomock.Any()).Return(repoGame, nil)
	srv = gamesrv.NewGameService(repo, generator)

	game, revealErr := srv.Revel("123", domain.Position{Col: 2, Row: 0})
	game.Debug()

	require.NoError(t, revealErr)
	require.NotNil(t, game)
	assert.Equal(t, game.Status, domain.InProgress)

	oneTip := domain.CellStatus(1)
	hidden := domain.Hidden

	assert.Equal(t, oneTip, game.Board[0][0])
	assert.Equal(t, oneTip, game.Board[0][1])
	assert.Equal(t, oneTip, game.Board[1][0])
	assert.Equal(t, oneTip, game.Board[0][2])
	assert.Equal(t, oneTip, game.Board[1][2])
	assert.Equal(t, hidden, game.Board[2][0])
	assert.Equal(t, hidden, game.Board[2][1])
	assert.Equal(t, hidden, game.Board[2][2])
	assert.Equal(t, domain.Mine, game.Board[1][1])

}

func TestRevelCell3(t *testing.T) {
	t.Parallel()
	var (
		srv  provides.GamesService
		ctrl *gomock.Controller
	)
	ctrl = gomock.NewController(t)
	ctrl.Finish()
	repo := mocks.NewMockGamesRepository(ctrl)
	generator := mocks.NewMockBoardGenerator(ctrl)
	repoGame := &domain.Game{
		Columns: 3,
		Lines:   3,
		Mines:   1,
	}
	board := domain.Board{
		{1, 1, 1},
		{1, -3, 1},
		{0, 0, 0},
	}
	repo.EXPECT().Get("123").Return(repoGame, nil)
	generator.EXPECT().New(3, 3, 1).Return(board, nil)
	repo.EXPECT().Save(gomock.Any()).Return(repoGame, nil)
	srv = gamesrv.NewGameService(repo, generator)

	game, revealErr := srv.Revel("123", domain.Position{Col: 1, Row: 2})
	game.Debug()

	require.NoError(t, revealErr)
	require.NotNil(t, game)
	assert.Equal(t, game.Status, domain.Win)

	oneTip := domain.CellStatus(1)

	assert.Equal(t, oneTip, game.Board[0][0])
	assert.Equal(t, oneTip, game.Board[0][1])
	assert.Equal(t, oneTip, game.Board[1][0])
	assert.Equal(t, oneTip, game.Board[0][2])
	assert.Equal(t, oneTip, game.Board[1][2])
	assert.Equal(t, oneTip, game.Board[2][0])
	assert.Equal(t, oneTip, game.Board[2][1])
	assert.Equal(t, oneTip, game.Board[2][2])
	assert.Equal(t, domain.Mine, game.Board[1][1])

}
