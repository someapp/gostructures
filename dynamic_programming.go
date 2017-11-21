package gostructures

// Given a maze and a robot that can only move right or down
// can the robot navigate from the top left to bottom right
// and if so return a possible path
func RobotMaze(maze [][]bool) ([]RobotMazePoint, bool) {
	// cache impossible paths
	cache := make(map[[2]int]bool)
	return robotMazePath(maze, len(maze)-1, len(maze[0])-1, &cache)
}

func robotMazePath(maze [][]bool, row int, col int, cache *map[[2]int]bool) ([]RobotMazePoint, bool) {
	point := RobotMazePoint{col, row}

	if row == 0 && col == 0 {
		return []RobotMazePoint{point}, true
	}

	if row < 0 || col < 0 || !maze[row][col] {
		return make([]RobotMazePoint, 0), false
	}

	// if we previously found this to be impossible
	// return false quickly
	cached, ok := (*cache)[point]
	if ok && !cached {
		return make([]RobotMazePoint, 0), false
	}

	if path, ok := robotMazePath(maze, row-1, col, cache); ok {
		path = append(path, point)
		return path, true
	}

	if path, ok := robotMazePath(maze, row, col-1, cache); ok {
		path = append(path, point)
		return path, true
	}

	(*cache)[point] = false
	return make([]RobotMazePoint, 0), false
}

type RobotMazePoint = [2]int
