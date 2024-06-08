package gamesrv

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/requires"
	"fmt"
)

func NewGameService(repository requires.GamesRepository,
	boardGenerator requires.BoardGenerator) *service {
	return &service{"", repository, boardGenerator}
}

type service struct {
	message        string
	repository     requires.GamesRepository
	boardGenerator requires.BoardGenerator
}

func (s *service) Error() string {
	return s.message
}

func (s *service) Create(lines, columns, mines int) (*domain.Game, error) {
	if lines < 1 {
		s.message = fmt.Sprintf("o tabuleiro do jogo deve ter um quantidade de linhas maior que um e recebeu %d", lines)

		return nil, s
	}

	if columns < 1 {
		s.message = fmt.Sprintf("deve ter pelo menos 2 colunas. Recebeu %d", columns)

		return nil, s
	}

	if mines < 1 {
		s.message = fmt.Sprintf("uma sessão de jogo deve possuir no mínimo uma mina e recebeu %d", mines)

		return nil, s
	}

	if mines >= lines*columns {
		s.message = fmt.Sprintf("uma sessão deve  ter menos do que %d minas e recebeu %d", lines*columns, mines)

		return nil, s
	}
	savedGame, err := s.repository.Save(domain.Game{
		Lines:   lines,
		Columns: columns,
		Mines:   mines,
	})
	if err != nil {
		return nil, err
	}
	return savedGame, nil
}

func (s *service) Get(id string) (*domain.Game, error) {
	if id == "" {
		return nil, fmt.Errorf(`id do jogo deve ser informado`)
	}
	game, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	if game != nil {
		return game, nil
	}

	return nil, fmt.Errorf(`jogo não existe`)
}

func (s *service) Revel(id string, pos domain.Position) (*domain.Game, error) {
	if id == "" {
		return nil, fmt.Errorf(`id do jogo deve ser informado`)
	}
	game, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, fmt.Errorf(`jogo não existe`)

	}
	if game.IsFinished() {
		// criar exibição
		return nil, fmt.Errorf(`jogo finalizado`)
	}

	// ### RN007 - Ao realizar a primeira revelação (jogo estado de novo)
	// o jogo de sortear a minas aleatoriamente ignorando a posição informada.
	if game.Status == domain.New {
		board, errB := s.boardGenerator.New(game.Lines, game.Columns, game.Mines)
		if errB != nil {
			return nil, fmt.Errorf(`erro ao criar tabuleiro`)
		}
		game.Board = board
		game.Status = domain.InProgress
	}

	if invalidPosition := game.InvalidPosition(pos); invalidPosition != nil {
		return nil, invalidPosition
	}

	if invalidCell := game.InvalidCell(pos); invalidCell != nil {
		return nil, invalidCell
	}

	if game.IsMine(pos) {
		game.Status = domain.Lost
	} else {
		game.Revel(pos)
	}

	if game.IsWin() {
		game.Status = domain.Win
	}

	// ### RN014 - O jogo deve ser gravado a cada alteração no estado do jogo ou do tabuleiro (revelar, marcar ou descmar celula)
	savedGame, errOnSave := s.repository.Save(*game)
	if errOnSave != nil {
		return nil, errOnSave
	}
	return savedGame, nil

}
