package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func getInput(s string) (id int, value *float64) {
	var v float64
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		panic(err)
	}
	str = strings.TrimSpace(str)
	input := strings.Split(str, ",")
	id, _ = strconv.ParseInt(input[0], 10, 64)
	if input[1] == "NULL" {
		value = nil
	} else {
		_, err := fmt.Sscanf(input[1], "%g",&v)
		if err != nil {
			panic(err)
		}
		value = &v
	}
	return id, value
}
