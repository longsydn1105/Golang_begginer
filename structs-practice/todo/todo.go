package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid Input")
	}

	return Todo{content}, nil
}

// func getUserInput(prompt string) string {
// 	fmt.Printf("%v ", prompt)

// 	reader := bufio.NewReader(os.Stdin)

// 	text, err := reader.ReadString('\n')

// 	if err != nil {
// 		fmt.Print("Invalid Input")
// 		return ""
// 	}

// 	return text
// }

// func getTodoData() string {
// 	return getUserInput("Totdo text: ")
// }
