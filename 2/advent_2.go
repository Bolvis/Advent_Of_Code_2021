package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input2.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fBody := string(f)

	fArr := strings.Split(fBody, "\n")

	var x int
	var y int

	for _, line := range fArr {
		lineCommandValue := strings.Fields(line)
		if len(lineCommandValue) == 0 {
			continue
		}
		value, _ := strconv.Atoi(lineCommandValue[1])
		if lineCommandValue[0] == "forward" {
			x += value
		}
		if lineCommandValue[0] == "up" {
			y -= value
		}
		if lineCommandValue[0] == "down" {
			y += value
		}
	}
	fmt.Println(x * y)
}
