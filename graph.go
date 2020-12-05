package graph

import (
	"fmt"
	"errors"
)

type Graph [][]interface{}

func (g *Graph) Cap () int {
	return cap(*g)
}

func (g *Graph) Len () int {
	return len(*g)
}

func (g *Graph) Set (row, col int, val interface{}) error {
	if row >= g.Len() || col >= len((*g)[row]) {
		s := fmt.Sprintf("Index (%d,%d) out of bounds (%dx%d)", row, col, g.Len(), len((*g)[row]))
		return errors.New(s)
	} else {
		(*g)[row][col] = val
	}
	return nil
}

func (g *Graph) Get (row, col int) (interface{}, error) {
	if row > g.Len() || col >= len((*g)[row]) {
		s := fmt.Sprintf("Index (%d,%d) out of bounds (%dx%d)", row, col, g.Len(), len((*g)[row]))
		return nil, errors.New(s)
	} else {
		return (*g)[row][col], nil
	}
}

func (g *Graph) Row (row int) ([]interface{}, error) {
	if row > g.Len() {
		s := fmt.Sprintf("Row (%d) out of bounds (%d)", row, g.Len())
		return nil, errors.New(s)
	} else {
		return (*g)[row], nil
	}
}

func (g *Graph) Col (col int) ([]interface{}, error) {
	var column []interface{}
	for row := 0; row < g.Len(); row++ {
		if col >= len((*g)[row]) {
			s := fmt.Sprintf("Col (%d) out of bounds on row %d (%d cols)", col, g.Len(), len((*g)[row]))
			return nil, errors.New(s)
		}
		column = append(column, (*g)[row][col])
	}
	return column, nil
}

func (g *Graph) Subgraph (r, c, w, h int) (*Graph, error) {
	var sub_g Graph = make([][]interface{}, h)
	for i := 0; i < h; i++ {
		sub_g[i] = make([]interface{}, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			src, err := g.Get(r+i,c+j)
			if err != nil {
				return nil, err
			}
			err = sub_g.Set(i, j, src)
			if err != nil {
				return nil, err
			}
		}
	}
	return &sub_g, nil
}

func (g *Graph) Map (f func(int, int, interface{})) {
	rows := g.Len()
	for row := 0; row < rows; row++ {
		cols := len((*g)[row])
		for col := 0; col < cols; col++ {
			f(row, col, (*g)[row][col])
		}
	}
}

func (g *Graph) Clone () *Graph {
	if nil == g {
		return nil
	}
	var clone Graph = make([][]interface{}, g.Len())
	for i := 0; i < g.Len(); i++ {
		clone[i] = make([]interface{}, len((*g)[i]))
		for j := 0; j < len((*g)[i]); j++ {
			clone[i][j] = (*g)[i][j]
		}
	}
	return &clone
}

func (g *Graph) String (f func(interface{}) string) string {
	s := "   "
	n := g.Len()
	for row := 0; row < n; row++ {
		s += fmt.Sprintf("  %d  ", row)
	}
	s += "\n"
	for row := 0; row < n; row++ {
		s += fmt.Sprintf("%d [", row)
		for col := 0; col < n; col++ {
			if (*g)[row][col] != nil {
				s += fmt.Sprintf(" %s ", f((*g)[row][col])) 
			} else {
				s += fmt.Sprintf("  _  ")
			}
		}
		s += "]\n"
	}
	return s
}