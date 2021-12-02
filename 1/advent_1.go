package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fBody := string(f)

	fArr := strings.Split(fBody, "\n")

	var result int32

	for index, line := range fArr {
		if index == 0 {
			continue
		}
		lineV, _ := strconv.Atoi(line)
		prevLineV, _ := strconv.Atoi(fArr[index-1])
		if lineV > prevLineV {
			result++
		}
	}
	fmt.Println(result)
}
