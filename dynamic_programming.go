package gostructures

// Given a maze and a robot that can only move right or down
// can the robot navigate from the top left to bottom right
// and if so return a possible path
func RobotMaze(maze [][]bool) ([]RobotMazePoint, bool) {
	return robotMazePath(maze, len(maze)-1, len(maze[0])-1)
}

func robotMazePath(maze [][]bool, row int, col int) ([]RobotMazePoint, bool) {
	if row == 0 && col == 0 {
		return []RobotMazePoint{{x: col, y: row}}, true
	}

	if row < 0 || col < 0 || !maze[row][col] {
		return make([]RobotMazePoint, 0), false
	}

	if path, ok := robotMazePath(maze, row-1, col); ok {
		path = append(path, RobotMazePoint{x: col, y: row})
		return path, true
	}

	if path, ok := robotMazePath(maze, row, col-1); ok {
		path = append(path, RobotMazePoint{x: col, y: row})
		return path, true
	}

	return make([]RobotMazePoint, 0), false
}

type RobotMazePoint struct {
	x int
	y int
}
