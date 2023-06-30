package main

import (
	"errors"
	"fmt"
)

func intersectSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("one of the slices is empty")
	}

	intersect := make([]int, 0)

	set := make(map[int]bool)
	for _, num := range slice1 {
		set[num] = true
	}

	for _, num := range slice2 {
		if set[num] {
			intersect = append(intersect, num)
		}
	}

	return intersect, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	intersect, err := intersectSlices(slice1, slice2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	fmt.Println("Intersection:", intersect)
}
