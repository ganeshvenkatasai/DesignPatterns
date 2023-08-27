package main

import "SnakesAndLadders/board"

func main() {
	board := &board.Board{}
	board = board.PlayerCount(2).Size(3).DiceCount(2).Snakes(
		map[int]int{
			10: 2,
			27: 8,
			99: 2,
		}).Ladders(
		map[int]int{
			6:  85,
			39: 52,
			65: 71,
		})
	board.Play()
}
