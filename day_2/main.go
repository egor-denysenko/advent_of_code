package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Game struct{
	Green int
	Red int
	Blue int
}

const red_max_value = 12
const green_max_value = 13
const blue_max_value = 14

func (g *Game) isGameValid() bool{
	if g.Red > red_max_value ||
	g.Green > green_max_value ||
	g.Blue > blue_max_value {return false};
	return true
}

func puzzle_1() {
	file, err := os.Open("./input")
	if err != nil{
		panic("error while reading the input file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	game_counter := 1
	result := 0
	games := make(map[int][]Game,0)
	for scanner.Scan(){
		games_values_inputs := strings.ReplaceAll(
			scanner.Text(),
			fmt.Sprintf("Game %d:",game_counter),
			"",
		);
		sub_games := strings.Split(games_values_inputs,";")		

		for _, sub_game := range sub_games {
			game := Game{}
			colors_input := strings.Split(sub_game, ",")
			for _, color := range colors_input {
				splitted_input := strings.Split(color[1:], " ")
				var val int
				if val, err = strconv.Atoi(splitted_input[0]); err != nil{
					continue
				}
				switch splitted_input[1]{
				case "green":
					game.Green += val
				case "red":
					game.Red += val
				case "blue":
					game.Blue += val
				}
			}
			games[game_counter] = append(games[game_counter], game)
		}
	game_counter++
	}

	for key, sub_game := range games {
		for _, game := range sub_game {
			if !game.isGameValid(){
				delete(games,key)
				break		
			}
		}
	}
	for key, _ := range games {
		result += key
	}
	fmt.Println("puzzle 1 result",result)
}

func main(){
	puzzle_1()
}

