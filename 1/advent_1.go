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
	f_body := string(f)

	f_arr := strings.Split(f_body, "\n")

	var result int32

	for index, line := range f_arr {
		if index == 0 {
			continue
		}
		line_v, _ := strconv.Atoi(line)
		prev_line_v, _ := strconv.Atoi(f_arr[index-1])
		if line_v > prev_line_v {
			result++
		}
	}
	fmt.Println(result)
}
