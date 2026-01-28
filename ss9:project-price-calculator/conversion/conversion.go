package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {

		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("ERROR: Failed to convert string to float.")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
