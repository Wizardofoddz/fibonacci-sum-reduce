package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	input := make([]uint64, 0)
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&input)
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println(Sum(input))
}

// Sum adds up all the numbers and returns
// the sum
func Sum(numbers []uint64) uint64 {
	sum := uint64(0)

	for _, number := range numbers {
		sum += number
	}

	return sum
}
