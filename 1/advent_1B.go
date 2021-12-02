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

	var result int

	for i := 0; i < len(fArr)-3; i++ {
		lineOneV, _ := strconv.Atoi(fArr[i])
		lineTwoV, _ := strconv.Atoi(fArr[i+1])
		lineThreeV, _ := strconv.Atoi(fArr[i+2])
		lineFourV, _ := strconv.Atoi(fArr[i+3])
		if lineOneV+lineTwoV+lineThreeV < lineTwoV+lineThreeV+lineFourV {
			result++
		}
	}
	fmt.Println(result)
}
