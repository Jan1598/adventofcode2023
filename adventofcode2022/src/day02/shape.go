package day02

import (
	"errors"
	"fmt"
)

type Shape struct {
	Letters      []string
	Name         string
	ChoosePoints int
}

type WinVariant struct {
	WinName   string
	LooseName string
}

const nameRock = "Rock"
const namePaper = "Paper"
const nameScissors = "Scissors"

var ShapeVariants = []Shape{{Letters: []string{"A", "X"}, Name: nameRock, ChoosePoints: 1},
	{Letters: []string{"B", "Y"}, Name: namePaper, ChoosePoints: 2},
	{Letters: []string{"C", "Z"}, Name: nameScissors, ChoosePoints: 3}}

var WinVariants = []WinVariant{{WinName: nameRock, LooseName: nameScissors},
	{WinName: nameScissors, LooseName: namePaper},
	{WinName: namePaper, LooseName: nameRock}}

func CheckPoints(letterToCheck, letterEnemy string) int {

	shapeToCheck, err := GetShapeByLetter(letterToCheck)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	enemyShape, err := GetShapeByLetter(letterEnemy)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	points := shapeToCheck.ChoosePoints
	if IsDraw(shapeToCheck, enemyShape) {
		points += 3
	} else if HasWon(shapeToCheck, enemyShape) {
		points += 6
	}
	return points
}

func GetShapeByLetter(letter string) (Shape, error) {
	for _, shape := range ShapeVariants {
		for _, letterInShape := range shape.Letters {
			if letterInShape == letter {
				return shape, nil
			}
		}
	}
	return Shape{}, errors.New("No shape found")
}

func HasWon(shapeToCheck, shapeEnemy Shape) bool {
	for _, winVariant := range WinVariants {
		if winVariant.WinName == shapeToCheck.Name && winVariant.LooseName == shapeEnemy.Name {
			return true
		}
	}
	return false
}

func GetCalcValueByLetter(letterToCheck, letterEnemy string) string {

	enemyShape, err := GetShapeByLetter(letterEnemy)
	if err != nil {
		fmt.Println(err)
		return letterToCheck
	}

	if letterToCheck == "Y" {
		return enemyShape.Letters[1]
	} else if letterToCheck == "X" {
		for _, shape := range ShapeVariants {
			if HasWon(enemyShape, shape) {
				return shape.Letters[1]
			}
		}
	} else if letterToCheck == "Z" {
		for _, shape := range ShapeVariants {
			if HasWon(shape, enemyShape) {
				return shape.Letters[1]
			}
		}
	}

	return letterToCheck
}

func IsDraw(shapeToCheck, shapeEnemy Shape) bool {
	return shapeToCheck.Name == shapeEnemy.Name
}
