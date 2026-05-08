package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	text, err := os.ReadFile("corpus/Metamorphosis.txt")
	if err != nil {
		log.Error("Error reading file: %v\n", err)
		return
	}
	stringText := string(text)
	regex := regexp.MustCompile(`(?m)^Chapter [IVXLCDM]+`)

	indices := regex.FindAllStringIndex(stringText, -1)
	var chapters []string
	for i, index := range indices {
		start := index[0]
		var end int
		if i+1 < len(indices) {
			end = indices[i+1][0]
		} else {
			end = len(stringText)
		}
		chapters = append(chapters, strings.TrimSpace(stringText[start:end]))
	}

	for i, chapter := range chapters {
		filename := fmt.Sprintf("corpus/chapters/Metamorphosis_Chapter_%d.txt", i+1)
		err = os.WriteFile(filename, []byte(chapter), 0644)
		if err != nil {
			log.Error("Error writing file: %v\n", err)
		}
	}

}
