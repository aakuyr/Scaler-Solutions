package main

func createDPArray(rowSize, colSize int) [][]int {
    DP := make([][]int, rowSize)
    for i := 0; i < rowSize; i++ {
        DP[i] = make([]int, colSize)
    }
    return DP
}

func uniquePathsWithObstacles(A [][]int )  (int) {
      rowSize := len(A)
      colSize := len(A[0])
      DP := createDPArray(rowSize, colSize)
      x := []int{0, -1}
      y := []int{-1, 0}
      return uniquePath(A, DP, rowSize - 1, colSize - 1, x, y)
}

func uniquePath(A, DP [][]int, n, m int, x, y []int) int {
    row := len(A)
    col := len(A[0])
    
    if n == 0 && m === 0 {
        DP[n][m] = 1
    }
    
    if DP[n][m] == 0 {
        for i := 0; i < 2; i++ {
            if n + x[i] >= 0 && n + x[i] < row &&
                m + y[i] >= 0 && m + y[i] < col && 
                A[n][m] != 1 {
                DP[n][m] += uniquePath(A, DP, n + x[i], m + y[i], x, y);                 
            }
        }
    }
    
    return DP[n][m];
}
