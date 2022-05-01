package main

import "fmt"

type Point struct {
	x int
	y int
}

func countRectangles(A, B []int) int {
	n := len(A)
	setOfPoints := make(map[Point]bool)
	for i := 0; i < len(A); i++ {
		p := Point{A[i], B[i]}
		if _, ok := setOfPoints[p]; ok == false {
			setOfPoints[p] = true
		}
	}

	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := Point{A[i], B[i]}
			b := Point{A[j], B[j]}

			p := Point{A[i], B[j]}
			q := Point{A[j], B[i]}

			if a.x == b.x || a.y == b.y {
				continue
			}
			_, pExists := setOfPoints[p]
			_, qExists := setOfPoints[q]

			if pExists == true && qExists == true {
				count++
			}
		}
	}

	return count / 2
}
func main() {
	A := []int{1, 1, 2, 2, 3, 3}
	B := []int{1, 2, 1, 2, 1, 2}
	fmt.Println(countRectangles(A, B))
}
