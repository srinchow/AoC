package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils"
	"os"
	"strings"
)

func getTurn(turn string) (string, string) {
	playerMoves := strings.Fields(turn)
	return playerMoves[0], playerMoves[1]
}

func part1(strategyGuide []string, playerMoveMap map[string]string, gameWinRules map[string]string) {
	totalScore := 0
	for _, turn := range strategyGuide {
		player1, player2 := getTurn(turn)
		baseScore := getBaseScore(playerMoveMap[player2])
		totalScore += baseScore
		totalScore += getTurnScore(playerMoveMap[player1], playerMoveMap[player2], gameWinRules)
	}

	fmt.Println(totalScore)
}

func part2(strategyGuide []string, playerMoveMap map[string]string, gameWinRules map[string]string) {
	totalScore := 0
	for _, turn := range strategyGuide {
		player1, expected := getTurn(turn)
		player1Move := playerMoveMap[player1]

		if expected == "X" { // lose case
			totalScore += getBaseScore(gameWinRules[gameWinRules[player1Move]])
		}

		if expected == "Y" {
			totalScore += getBaseScore(player1Move)
			totalScore += 3
		}

		if expected == "Z" {
			totalScore += getBaseScore(gameWinRules[player1Move])
			totalScore += 6
		}

	}
	fmt.Println(totalScore)

}

func main() {
	file, err := os.Open("./Day2/input.txt")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
	}
	defer utils.CloseFile(file)

	moves := utils.ParseFile(file)
	rockPaperScissor := cretePlayerMoveMap()
	gameRules := createGameRulesWinMap()

	if len(moves) > 1 {
		fmt.Println("Error parsing file")
		return
	}

	part1(moves[0], rockPaperScissor, gameRules)
	part2(moves[0], rockPaperScissor, gameRules)

}

func createGameRulesWinMap() map[string]string {
	winningMap := make(map[string]string)
	winningMap["ROCK"] = "PAPER"
	winningMap["PAPER"] = "SCI"
	winningMap["SCI"] = "ROCK"
	return winningMap
}

func cretePlayerMoveMap() map[string]string {
	moveMapping := make(map[string]string)
	moveMapping["A"] = "ROCK"
	moveMapping["X"] = "ROCK"

	moveMapping["B"] = "PAPER"
	moveMapping["Y"] = "PAPER"

	moveMapping["C"] = "SCI"
	moveMapping["Z"] = "SCI"

	return moveMapping

}

func getBaseScore(player2 string) int {
	switch player2 {
	case "ROCK":
		return 1
	case "PAPER":
		return 2
	case "SCI":
		return 3
	default:
		return 0
	}
}

func getTurnScore(p1, p2 string, rules map[string]string) int {
	if p1 == p2 {
		return 3
	}
	if rules[p1] == p2 {
		return 6
	}
	return 0
}
