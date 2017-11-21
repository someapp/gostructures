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
	// -X-X
	// -X-X
	// X---
	maze := [][]bool{
		[]bool{true, false, true, true},
		[]bool{true, false, true, false},
		[]bool{true, false, true, false},
		[]bool{false, true, true, true},
	}

	_, ok := RobotMaze(maze)
	if ok {
		t.Error("maze should not be completable")
	}
}
