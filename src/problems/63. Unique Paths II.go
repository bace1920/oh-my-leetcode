package main

import "fmt"

func main() {
	// obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	// obstacleGrid := [][]int{{0, 1}, {0, 0}}
	// obstacleGrid := [][]int{{0, 0}, {0, 1}}
	obstacleGrid := [][]int{{1}, }
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])
	cur := make([]int, width)

	if obstacleGrid[0][0] == 1 {
		cur[0] = -1
	} else {

		cur[0] = 1
	}
	for q := 1; q < width; q++ {
		if obstacleGrid[0][q] == 1 {
			cur[q] = -1
		} else {
			cur[q] = cur[q-1]
		}
	}

	for i := 1; i < height; i++ {
		if obstacleGrid[i][0] == 1 {
			cur[0] = -1
		}
		for q := 1; q < width; q++ {
			// if cur[q-1] == -1 && cur[q] != -1 {
			// 	cur[q] = cur[q]
			// } else if cur[q-1] == -1 && cur[q] == -1 {
			// 	cur[q] = cur[q]
			// }
			if obstacleGrid[i][q] == 1 {
				cur[q] = -1
			} else if cur[q-1] != -1 && cur[q] != -1 {
				cur[q] += cur[q-1]
			} else if cur[q-1] != -1 && cur[q] == -1 {
				cur[q] = cur[q-1]
			}
		}
	}
	if cur[width-1] < 0 {
		return 0
	}
	return cur[width-1]
}
