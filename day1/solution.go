package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func run() {
	arrs := readInput()
	fmt.Printf("day1 part1 result: %d\n", part1(arrs[0], arrs[1]))
	fmt.Printf("day1 part2 result: %d\n", part2(arrs[0], arrs[1]))
}

func part1(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += int(math.Abs(float64(arr1[i] - arr2[i])))
	}
	return sum
}

func part2(arr1, arr2 []int) int {
	appears := make(map[rune]int, len(arr1))
	for i := 0; i < len(arr2); i++ {
		foundVal, ok := appears[rune(arr2[i])]
		if !ok {
			appears[rune(arr2[i])] = 1
		} else {
			appears[rune(arr2[i])] = foundVal + 1
		}
	}
	sum := 0
	for i := 0; i < len(arr1); i++ {
		foundVal, ok := appears[rune(arr1[i])]
		if ok {
			sum += arr1[i] * foundVal
		}
	}
	return sum
}

func readInput() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	var arr1 []int
	var arr2 []int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		nums := strings.Split(text, "   ")
		num, _ := strconv.Atoi(nums[0])
		arr1 = append(arr1, num)
		num, _ = strconv.Atoi(nums[1])
		arr2 = append(arr2, num)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return [][]int{arr1, arr2}
}
