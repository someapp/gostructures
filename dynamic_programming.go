package gostructures

// Given a maze and a robot that can only move right or down
// can the robot navigate from the top left to bottom right
// and if so return a possible path
func RobotMaze(maze [][]bool) ([]RobotMazePoint, bool) {
	cache := make(map[[2]int]bool)
	return robotMazePath(maze, len(maze)-1, len(maze[0])-1, &cache)
}

func robotMazePath(maze [][]bool, row int, col int, cache *map[[2]int]bool) ([]RobotMazePoint, bool) {
	if row == 0 && col == 0 {
		return []RobotMazePoint{{x: col, y: row}}, true
	}

	if row < 0 || col < 0 || !maze[row][col] {
		return make([]RobotMazePoint, 0), false
	}

	cached, ok := (*cache)[[2]int{col, row}]
	if ok && !cached {
		return make([]RobotMazePoint, 0), false
	}

	if path, ok := robotMazePath(maze, row-1, col, cache); ok {
		path = append(path, RobotMazePoint{x: col, y: row})
		return path, true
	}

	if path, ok := robotMazePath(maze, row, col-1, cache); ok {
		path = append(path, RobotMazePoint{x: col, y: row})
		return path, true
	}

	(*cache)[[2]int{col, row}] = false
	return make([]RobotMazePoint, 0), false
}

type RobotMazePoint struct {
	x int
	y int
}
