package gostructures

// Given a maze and a robot that can only move right or down
// can the robot navigate from the top left to bottom right
// and if so return a possible path
func RobotMaze(maze [][]bool) ([]RobotMazePoint, bool) {
	// cache unreachable points
	unreachable := make(map[RobotMazePoint]bool)
	return robotMazePath(maze, len(maze)-1, len(maze[0])-1, &unreachable)
}

func robotMazePath(maze [][]bool, row int, col int, unreachable *map[RobotMazePoint]bool) ([]RobotMazePoint, bool) {
	point := RobotMazePoint{col, row}

	if row == 0 && col == 0 {
		return []RobotMazePoint{point}, true
	}

	if row < 0 || col < 0 || !maze[row][col] {
		return make([]RobotMazePoint, 0), false
	}

	// if we previously found this to be impossible
	// return false quickly
	if (*unreachable)[point] {
		return make([]RobotMazePoint, 0), false
	}

	if path, ok := robotMazePath(maze, row-1, col, unreachable); ok {
		path = append(path, point)
		return path, true
	}

	if path, ok := robotMazePath(maze, row, col-1, unreachable); ok {
		path = append(path, point)
		return path, true
	}

	(*unreachable)[point] = true
	return make([]RobotMazePoint, 0), false
}

type RobotMazePoint = [2]int
