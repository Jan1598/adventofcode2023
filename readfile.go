package adventofcode2023

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadData(filePath string) []string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Fehler beim Abrufen des aktuellen Arbeitsverzeichnisses:", err)
	}
	dat, err := os.ReadFile(filepath.Join(currentDir, filePath))
	check(err)
	return strings.Split(strings.ReplaceAll(string(dat), "\r", ""), "\n")
}
