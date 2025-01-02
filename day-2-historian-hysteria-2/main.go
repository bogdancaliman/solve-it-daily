package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	var list1, list2 []int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %v", line)
		}

		num1 := strToInt(parts[0])
		num2 := strToInt(parts[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	score := 0
	for i := 0; i < len(list1); i++ {
		count := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				count++
			}
		}

		score += count * list1[i]
	}

	fmt.Println("Score:", score)

}

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting %s to int: %v", s, err)
	}
	return num
}
