package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	ROWS = 9
	COLUMNS = 9

    EMPTY = "□"
    BLOCK = "■"

    LEFT = "⇦"
    RIGHT = "⇨"
    UP = "⇧"
    DOWN = "⇩"
)

var visit[ROWS][COLUMNS] bool
var mazeStack []XY
var reversePath []XY
var maze [ROWS][COLUMNS] string

type XY struct   {
	x int
	y int
}

func isValidXY(xy XY) bool {
	x := xy.x
	y := xy.y
	if x < 0 || x >= COLUMNS || y < 0 || y >= ROWS {
		return false
	}

	return true
}

func isDead(xy XY) bool {
	left := XY{xy.x-1,xy.y}
	right := XY{xy.x+1,xy.y}
	up := XY{xy.x,xy.y-1}
	down := XY{xy.x,xy.y+1}

	if isValidXY(left) && maze[left.y][left.x] == EMPTY && !visit[left.y][left.x] ||
		isValidXY(right) && maze[right.y][right.x] == EMPTY &&!visit[right.y][right.x]||
		isValidXY(up) && maze[up.y][up.x] == EMPTY &&!visit[up.y][up.x]||
		isValidXY(down) && maze[down.y][down.x] == EMPTY&&!visit[down.y][down.x] {
		return false
	}

	return true
}

func mazeHandle() {
	if maze[0][0] != EMPTY {
		return
	}

	mazeStack = append(mazeStack, XY{0, 0})

	for ; len(mazeStack) > 0 ; {
		top := mazeStack[len(mazeStack)-1]
		if top.x == COLUMNS - 1 && top.y == ROWS - 1 {

			reversePath = append(reversePath, top)
			for ; len(mazeStack) > 0 ; {
				top = mazeStack[len(mazeStack)-1]
				prev := XY{reversePath[len(reversePath)-1].x, reversePath[len(reversePath)-1].y}
				if top.x + top.y - prev.x - prev.y == 1 ||
					top.x + top.y - prev.x - prev.y == -1 {
					reversePath = append(reversePath, top)
				}
				mazeStack = mazeStack[:len(mazeStack) - 1]
			}
			var path []XY
			for i := len(reversePath) - 1;i>=0;i-- {
				path = append(path, reversePath[i])
			}
			for i:= 1;i< len(path);i++ {
				prev := path[i-1]
				cur := path[i]
				if cur.x - prev.x == 0 {
					if cur.y - prev.y == 1 {
						maze[cur.y][cur.x] = DOWN
					}else {
						maze[cur.y][cur.x] = UP
					}

				} else if cur.x - prev.x == 1 {
					maze[cur.y][cur.x] = RIGHT
				}else{
					maze[cur.y][cur.x] = LEFT
				}
			}

			return
		}

		if isDead(top) {
			visit[top.y][top.x] = true
			mazeStack = mazeStack[:len(mazeStack) - 1]
			continue
		}

		visit[top.y][top.x] = true

		down := XY{top.x,top.y+1}
		if isValidXY(down) && maze[down.y][down.x] == EMPTY && !visit[down.y][down.x]{
			mazeStack = append(mazeStack, down)
		}
		right := XY{top.x+1,top.y}
		if isValidXY(right) && maze[right.y][right.x] == EMPTY && !visit[right.y][right.x]{
			mazeStack = append(mazeStack, right)
		}
		left := XY{top.x-1,top.y}
		if isValidXY(left) && maze[left.y][left.x] == EMPTY && !visit[left.y][left.x] {
			mazeStack = append(mazeStack, left)
		}
		up := XY{top.x,top.y-1}
		if isValidXY(up) && maze[up.y][up.x] == EMPTY && !visit[up.y][up.x]{
			mazeStack = append(mazeStack, up)
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < ROWS; i++ {
		for j:= 0; j < COLUMNS; j++ {
			if rand.Int() % 3 == 0 {
				maze[i][j] = BLOCK
			}else {
				maze[i][j] = EMPTY
			}
		}
	}

	mazeHandle()

	for i := 0; i < ROWS; i++ {
		for j:= 0; j < COLUMNS; j++ {
			fmt.Print(maze[i][j])
		}
		fmt.Println()
	}
}








