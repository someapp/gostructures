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
		{x: 0, y: 0},
		{x: 0, y: 1},
		{x: 1, y: 1},
		{x: 2, y: 1},
		{x: 2, y: 2},
		{x: 2, y: 3},
		{x: 3, y: 3},
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
