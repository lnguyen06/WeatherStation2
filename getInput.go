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
	id, _ = strconv.Atoi(input[0])
	if input[1] == "NULL" {
		value = nil
	} else {
		_, err  := fmt.Sscanf(input[1], "%g",&v)
		if err != nil {
			panic(err)
		}
		value = &v
	}
	return id, value
}

func main() {
	s := "11,15.5"
	id, value := getInput(s)
	fmt.Println("id, value: ", id, value)
}