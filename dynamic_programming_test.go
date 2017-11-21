package gostructures

import (
	"reflect"
	"testing"
)

func TestRobotMaze_Completable(t *testing.T) {
	// -X--
	// ---X
	// -X-X
	// X---
	maze := [][]bool{
		[]bool{true, false, true, true},
		[]bool{true, true, true, false},
		[]bool{true, false, true, false},
		[]bool{false, true, true, true},
	}

	path, ok := RobotMaze(maze)
	if !ok {
		t.Error("maze should be completable")
	}

	expectedPath := []RobotMazePoint{
		{0, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{2, 2},
		{2, 3},
		{3, 3},
	}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Error("unexpected path")
	}
}

func TestRobotMaze_NotCompletable(t *testing.T) {
	// -X--
	// X---
	// ----
	// ----
	maze := [][]bool{
		[]bool{true, false, true, true},
		[]bool{false, true, true, true},
		[]bool{true, true, true, true},
		[]bool{true, true, true, true},
	}

	_, ok := RobotMaze(maze)
	if ok {
		t.Error("maze should not be completable")
	}
}

func TestRobotMaze_Large(t *testing.T) {
	// a massive empty maze
	// --...--
	// .......
	// --...--
	maze := make([][]bool, 1000)

	for i := 0; i < 1000; i++ {
		maze[i] = make([]bool, 1000)
		for j := 0; j < 1000; j++ {
			maze[i][j] = true
		}
	}

	path, ok := RobotMaze(maze)
	if !ok {
		t.Error("maze should be completable")
	}

	expectedPath := make([]RobotMazePoint, 0, 1999)
	for i := 0; i < 1000; i++ {
		expectedPath = append(expectedPath, RobotMazePoint{i, 0})
	}

	for j := 1; j < 1000; j++ {
		expectedPath = append(expectedPath, RobotMazePoint{999, j})
	}

	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("unexpected path %v", path)
	}
}

func TestRobotMaze_Large_Fail(t *testing.T) {
	// a massive empty maze
	// that looks solvable
	// but isnt
	// -X...--
	// X-...--
	// .......
	// --...--
	maze := make([][]bool, 1000)

	for i := 0; i < 1000; i++ {
		maze[i] = make([]bool, 1000)
		for j := 0; j < 1000; j++ {
			maze[i][j] = true
		}
	}

	maze[1][0] = false
	maze[0][1] = false

	_, ok := RobotMaze(maze)
	if ok {
		t.Error("maze should not be completable")
	}
}
