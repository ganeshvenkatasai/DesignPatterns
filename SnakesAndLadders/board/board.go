package board

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SIX = 6
)

type Board struct {
	playerCount int
	boardSize   int
	diceCount   int
	snakes      map[int]int
	ladders     map[int]int
}

func (b *Board) PlayerCount(n int) *Board {
	if n < 1 {
		n = 1
	}
	b.playerCount = n
	return b
}

func (b *Board) Size(s int) *Board {
	if s < 1 {
		s = 10
	}
	b.boardSize = s * s
	return b
}

func (b *Board) DiceCount(c int) *Board {
	if c < 1 {
		c = 1
	}
	b.diceCount = c
	return b
}

func (b *Board) Snakes(s map[int]int) *Board {
	for k, v := range s {
		if k <= v || k == b.boardSize {
			delete(s, k)
		}
	}
	b.snakes = s
	return b
}

func (b *Board) Ladders(l map[int]int) *Board {
	for k, v := range l {
		if k >= v {
			delete(l, k)
		}
	}
	b.ladders = l
	return b
}

func (b *Board) Play() {
	ind := 0
	player := make([]int, b.playerCount)
	for i := 0; i < len(player); i++ {
		player[i] = i + 1
	}
	playerPosition := make([]int, b.playerCount+1)
	for len(player) > 0 {
		currentPlayer := player[ind]
		fmt.Println("Player", currentPlayer, "turn")
		time.Sleep(time.Second)
		val := rand.Intn(SIX*b.diceCount-b.diceCount) + b.diceCount
		if dest := playerPosition[currentPlayer] + val; dest <= b.boardSize {
			playerPosition[currentPlayer] += val
		}
		fmt.Println("Dice value is", val)
		fmt.Println("Player", currentPlayer, "moved to", playerPosition[currentPlayer])
		if playerPosition[currentPlayer] == b.boardSize {
			fmt.Println("Player", currentPlayer, "got", b.playerCount-len(player)+1, "position")
			player = append(player[:currentPlayer], player[currentPlayer+1:]...)
		} else {
			ind++
		}
		if ind >= len(player) {
			ind = 0
		}
	}
}
