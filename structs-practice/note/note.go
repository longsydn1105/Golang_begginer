package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Println("-------------------")
	fmt.Printf("Your note title: %v \n\n", note.Title)
	fmt.Printf("Has content: %v \n\n", note.Content)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Create Note fail!!")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
