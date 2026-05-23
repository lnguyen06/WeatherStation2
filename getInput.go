package main
import (
	"fmt"
	"string"
	"bufio"
	"os"
	"strconv"
)

func getInput() (int, *float64) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	input = strings.Split(input, ',')
	id := 
}