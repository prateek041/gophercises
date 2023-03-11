package page

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

type Stories struct {
	StoryID Story `json:"story_id"`
}

type Story struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func StartGame() {
	storySlice := ReadAll()               // read all the data from stories file.
	stories := createStoryMap(storySlice) // create a story map from the read data
	_ = stories
}

func createStoryMap(f []byte) []Stories {
	var stories []Stories
	err := json.Unmarshal(f, &stories)

	if err != nil {
		log.Fatal("Error in unmarshalling the data", err)
	}

	fmt.Printf("%v \n", stories)
	return stories
}

func ReadAll() []byte {
	// we will use unmarshal here
	pwd, _ := os.Getwd()
	filePath := path.Dir(pwd) + "/gopher.json"
	// Read the file
	f, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("error reading file", err)
	}
	return f
}
