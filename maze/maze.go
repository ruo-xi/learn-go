package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var row, col int

	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	fmt.Println(row, col)
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(gird [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(gird) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(gird[0]) {
		return 0, false
	}
	return gird[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)

			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		println()
	}

	steps := walk(maze, point{0, 0}, point{2, 2})

	println()
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		println()
	}
}
