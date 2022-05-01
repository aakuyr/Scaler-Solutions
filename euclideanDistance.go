package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x int
	y int
}

type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	sumOfI := math.Sqrt(float64(p[i].x*p[i].x + p[i].y*p[i].y))
	sumOfJ := math.Sqrt(float64(p[j].x*p[j].x + p[j].y*p[j].y))
	return sumOfI < sumOfJ
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getBClosestEuclideanPoints(A [][]int, B int) [][]int {
	n := len(A)
	pointsArray := make([]Point, n)

	for i := 0; i < n; i++ {
		pointsArray[i] = Point{A[i][0], A[i][1]}
	}
	sort.Sort(Points(pointsArray))

	result := make([][]int, B)
	for i := 0; i < B; i++ {
		result[i] = []int{pointsArray[i].x, pointsArray[i].y}
	}

	return result
}

func main() {
	A := [][]int{
		{1, 3},
		{-2, 2},
	}
	fmt.Println(getBClosestEuclideanPoints(A, 2))
}
