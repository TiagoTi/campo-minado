package domain

type Position struct {
	Col int
	Row int
}

type Positions []Position

func (p Position) Adjcts() Positions {
	return Positions{
		Position{Row: p.Row - 1, Col: p.Col + 1}, //inf dir
		Position{Row: p.Row - 1, Col: p.Col - 1}, //inf esq
		Position{Row: p.Row + 1, Col: p.Col + 1}, //sup dir
		Position{Row: p.Row + 1, Col: p.Col - 1}, //sup esq
		Position{Row: p.Row, Col: p.Col + 1},     // dir
		Position{Row: p.Row, Col: p.Col - 1},     // esq
		Position{Row: p.Row + 1, Col: p.Col},     //sup
		Position{Row: p.Row - 1, Col: p.Col},     //inf
	}
}

// func (p Position) Igual(o Position) bool {
// 	return o.Row == p.Row && o.Col == p.Col
// }

// func (ps Positions) Exists(o Position) bool {
// 	for _, p := range ps {
// 		if p.Row == o.Row && p.Col == o.Col {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (posicoes Posicoes) String() string {
// 	s := fmt.Sprintf("len:%d [", len(posicoes))
// 	for _, p := range posicoes {
// 		s = fmt.Sprintf("%s %s,", s, p)
// 	}
// 	s = fmt.Sprintf("%s ]", s)
// 	return s
// }

// func (p Position) String() string {
// 	return fmt.Sprintf("(x:%d,y:%d)", p.Col, p.Row)
// }
