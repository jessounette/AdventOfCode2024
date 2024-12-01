package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 - distanceList : ", distanceList())
}

func distanceList() int {
	var list1 []int
	var list2 []int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			result := strings.Fields(line)

			var elem1, _ = strconv.Atoi(result[0])
			var elem2, _ = strconv.Atoi(result[1])

			list1 = append(list1, elem1)
			list2 = append(list2, elem2)
		}
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	var distanceList = 0

	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			distanceList += list1[i] - list2[i]
		} else {
			distanceList += list2[i] - list1[i]
		}
	}
	return distanceList
}
