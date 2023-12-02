package day02

import (
	"strconv"
	"strings"
)

type GameData struct {
	Id    int
	Cubes []GameCubes
}

type GameCubes struct {
	BlueCubes  int
	RedCubes   int
	GreenCubes int
}

func CubeDataFromLine(line string) GameData {

	if !strings.Contains(line, "Game") || !strings.Contains(line, ":") {
		return GameData{}
	}

	cubeText := strings.Split(line, ":")
	idText := strings.Trim(strings.Replace(cubeText[0], "Game", "", -1), " ")
	id, _ := strconv.Atoi(idText)

	cubeSet := strings.Split(cubeText[1], ";")

	var gameCubes []GameCubes
	gameData := GameData{
		Id:    id,
		Cubes: gameCubes,
	}

	for _, cube := range cubeSet {
		gameData.CollectGameData(cube)
	}

	return gameData
}

func (gameData *GameData) CollectGameData(cubeSet string) {

	var colorSet []string
	if strings.Contains(cubeSet, ",") {
		colorSet = strings.Split(cubeSet, ",")
	} else {
		colorSet = append(colorSet, cubeSet)
	}

	cubes := GameCubes{
		BlueCubes:  0,
		RedCubes:   0,
		GreenCubes: 0,
	}

	for _, colorText := range colorSet {
		cubes.CollectCubes(colorText)
	}

	gameData.Cubes = append(gameData.Cubes, cubes)
}

func (gameCubes *GameCubes) CollectCubes(colorText string) {
	if strings.Contains(colorText, "blue") {
		gameCubes.BlueCubes = gameCubes.BlueCubes + GetNumberFromColorText(colorText, "blue")
	} else if strings.Contains(colorText, "red") {
		gameCubes.RedCubes = gameCubes.RedCubes + GetNumberFromColorText(colorText, "red")
	} else if strings.Contains(colorText, "green") {
		gameCubes.GreenCubes = gameCubes.GreenCubes + GetNumberFromColorText(colorText, "green")
	}
}

func GetNumberFromColorText(colorText, color string) int {
	num, err := strconv.Atoi(strings.Trim(strings.Replace(colorText, color, "", -1), " "))
	if err != nil {
		return 0
	}
	return num
}
