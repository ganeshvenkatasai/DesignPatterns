package main

import "fmt"

type Board struct {
	Size  int
	Count int
}

func NewBoard(n int) *Board {
	return &Board{Size: n, Count: n * n}
}

func (b *Board) Move(m [][]int, r int, c int, p int) {
	m[r][c] = p
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i], " ")
		}
		fmt.Println()
	}
}

func main() {
	var row, col, n, player int
	fmt.Scanln(&n)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	b := NewBoard(n)
	for b.Count > 0 {
		player = 1
		fmt.Println("Player ", player, " : ")
		fmt.Scanln(&row, &col)
		b.Move(matrix, row, col, player)
		player = 2
		fmt.Println("Player ", player, " : ")
		fmt.Scanln(&row, &col)
		b.Move(matrix, row, col, player)
	}
}
