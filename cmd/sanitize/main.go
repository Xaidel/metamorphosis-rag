package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

type Metadata struct {
	Book      string `json:"book"`
	Chapter   string `json:"chapter"`
	Paragraph int    `json:"paragraph"`
	Source    string `json:"source"`
}

type Content struct {
	ID       string   `json:"id"`
	Text     string   `json:"text"`
	Metadata Metadata `json:"metadata"`
}

func main() {
	text, err := extractStory("corpus/Metamorphosis.txt")
	if err != nil {
		log.Error("Error extracting story: %v\n", err)
		return
	}
	detectChapterBoundaries(text)
	chap1file := "corpus/chapters/1/Metamorphosis_Chapter_1.txt"
	chap1txt, err := os.ReadFile(chap1file)
	chap2file := "corpus/chapters/2/Metamorphosis_Chapter_2.txt"
	chap2txt, err := os.ReadFile(chap2file)
	chap3file := "corpus/chapters/3/Metamorphosis_Chapter_3.txt"
	chap3txt, err := os.ReadFile(chap3file)
	if err != nil {
		log.Error("Error reading chapter file: %v\n", err)
		return
	}
	paragraphs1 := rebuildParagraphs(string(chap1txt))
	paragraphs2 := rebuildParagraphs(string(chap2txt))
	paragraphs3 := rebuildParagraphs(string(chap3txt))
	convertToMetadata(paragraphs1, 1)
	convertToMetadata(paragraphs2, 2)
	convertToMetadata(paragraphs3, 3)

}

func convertToMetadata(paragraphs []string, chapterNum int) {
	for i, p := range paragraphs {
		metadata := Metadata{
			Book:      "Metamorphosis",
			Chapter:   fmt.Sprintf("%d", chapterNum),
			Paragraph: i + 1,
			Source:    fmt.Sprintf("corpus/chapters/%d/Metamorphosis_Chapter_%d.txt", chapterNum, chapterNum),
		}
		content := Content{
			ID:       fmt.Sprintf("metamorphosis_chapter_%d_paragraph_%d", chapterNum, i+1),
			Text:     p,
			Metadata: metadata,
		}

		data, err := json.Marshal(content)
		if err != nil {
			log.Error("Error marshaling content to JSON: %v\n", err)
			continue
		}

		err = os.WriteFile(fmt.Sprintf("corpus/chapters/%d/Metamorphosis_Chapter_%d_Paragraph_%d.json", chapterNum, chapterNum, i+1), data, 0644)
		if err != nil {
			log.Error("Error writing JSON file: %v\n", err)
		}
	}
}

func rebuildParagraphs(text string) []string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	paragraphs := strings.Split(text, "\n\n")
	fmt.Println(len(paragraphs))
	var cleanedParagraphs []string
	for _, p := range paragraphs {
		cleaned := strings.ReplaceAll(p, "\n", " ")
		cleanedParagraphs = append(cleanedParagraphs, cleaned)
	}
	return cleanedParagraphs
}

func extractStory(path string) (string, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		log.Error("Error reading file: %v\n", err)
		return "", err
	}
	return string(text), nil
}

func detectChapterBoundaries(text string) {
	regex := regexp.MustCompile(`(?m)^Chapter [IVXLCDM]+`)
	if err := os.MkdirAll("corpus/chapters", 0755); err != nil {
		log.Error("Error creating chapters directory: %v\n", err)
		return
	}

	indices := regex.FindAllStringIndex(text, -1)
	var chapters []string
	for i, index := range indices {
		start := index[1]
		var end int
		if i+1 < len(indices) {
			end = indices[i+1][0]
		} else {
			end = len(text)
		}
		chapters = append(chapters, strings.TrimSpace(text[start:end]))
	}

	for i, chapter := range chapters {
		dir := fmt.Sprintf("corpus/chapters/%d", i+1)
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Error("Error creating directory: %v\n", err)
			continue
		}
		filename := fmt.Sprintf("corpus/chapters/%d/Metamorphosis_Chapter_%d.txt", i+1, i+1)
		err := os.WriteFile(filename, []byte(chapter), 0644)
		if err != nil {
			log.Error("Error writing file: %v\n", err)
		}
	}
}
