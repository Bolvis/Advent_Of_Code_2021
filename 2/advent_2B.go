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
	f_body := string(f)

	f_arr := strings.Split(f_body, "\n")

	var x int
	var y int
	var aim int

	for _, line := range f_arr {
		line_command_value := strings.Fields(line)
		if len(line_command_value) == 0 {
			continue
		}
		value, _ := strconv.Atoi(line_command_value[1])
		if line_command_value[0] == "forward" {
			x += value
			y += aim * value
		}
		if line_command_value[0] == "up" {
			aim -= value
		}
		if line_command_value[0] == "down" {
			aim += value
		}
	}
	fmt.Println(x * y)
}
