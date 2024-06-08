package domain

import "fmt"

type GameStatus int

const (
	New GameStatus = iota
	InProgress
	Lost
	Win
)

type Game struct {
	ID      string
	Lines   int
	Columns int
	Mines   int
	Board   Board
	Status  GameStatus
}

func (g *Game) Revel(pos Position) {
	if g.invalidPosition(pos) || g.isMine(pos) || g.isTip(pos) || g.isReveled(pos) {
		return
	}

	g.Reveled(pos)
	g.countAdjMines(pos)

	for _, p := range pos.Adjcts() {
		if g.invalidPosition(p) || g.isMine(p) || g.isTip(p) || g.isReveled(p) {
			continue
		}
		if g.isTip(pos) {
			g.countAdjMines(p)
		} else {
			g.Revel(p)
		}
	}
}

func (g *Game) countAdjMines(p Position) {
	adjctsMines := 0
	for _, adj := range p.Adjcts() {
		if g.invalidPosition(adj) {
			continue
		}
		if g.isMine(adj) {
			adjctsMines++
		}
	}
	if adjctsMines > 0 {
		g.setTip(p, CellStatus(adjctsMines))
	}
}
func (g *Game) setTip(p Position, mines CellStatus) {
	g.Board[p.Row][p.Col] = mines
}
func (g *Game) Reveled(p Position) {
	g.Board[p.Row][p.Col] = Reveled
}
func (g *Game) invalidPosition(p Position) bool {
	return p.Row < 0 || p.Col < 0 || p.Row >= len(g.Board) || p.Col >= len(g.Board[0])
}
func (g *Game) isMine(pos Position) bool {
	return g.Board[pos.Row][pos.Col] == Mine
}
func (g *Game) isReveled(pos Position) bool {
	return g.Board[pos.Row][pos.Col] == Reveled
}
func (g *Game) IsMine(pos Position) bool {
	return g.isMine(pos)
}

func (g *Game) isTip(pos Position) bool {
	return g.Board[pos.Row][pos.Col] > 0
}

func (g *Game) IsFinished() bool {
	if g.Status == Lost || g.Status == Win {
		return true
	}

	return false
}

// invalidPosition -> RN005: O jogador só pode realizar revelações em posições válidas do tabuleiro.
func (g *Game) InvalidPosition(p Position) error {
	if p.Row < 0 {
		return fmt.Errorf("não existe linha menor que zero")
	}
	if p.Col < 0 {
		return fmt.Errorf("não existe coluna menor que zero")
	}
	if p.Row >= len(g.Board) {
		return fmt.Errorf("esta linha não exite no tabuleiro")
	}
	if p.Col >= len(g.Board[0]) {
		return fmt.Errorf("esta coluna não existe no tabuleiro")
	}
	return nil
}

// invalidCell -> RN006 - O jogador não pode revelar celular abertas, e não pode revelar células com dicas
func (g *Game) InvalidCell(p Position) error {
	if g.isReveled(p) {
		return fmt.Errorf("esta celula já foi revelada")
	}
	if g.isTip(p) {
		return fmt.Errorf("esta celula é uma dica e já foi revelada")
	}
	return nil
}

func (g *Game) Debug() {
	for _, l := range g.Board {
		fmt.Println(l)
	}
}
func (g *Game) IsWin() bool {
	r := 0
	for _, l := range g.Board {
		for _, c := range l {
			if c == Reveled || c>0{
				r++
			}
		}
	}
	cells := (g.Columns * g.Lines) - g.Mines

	return cells == r
}
