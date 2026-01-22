package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile(balanceFileName string) (float64, error) {
	data, err := os.ReadFile(balanceFileName)

	if err != nil {
		return 1000, errors.New("Fail to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Fail to convert balance")
	}
	return balance, nil
}

func WriteBalanceToFile(balanceFileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}
