package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"math/rand"
	"time"
)


const width, height = int32(1240), int32(1240)
const mazeWidth, mazeHeight = 31, 31

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	
	size := int32(40)
	
	var maze [mazeWidth][mazeHeight] bool
	var tempMaze [mazeWidth][mazeWidth] int

	for i := int32(0); i < mazeWidth; i++ {
		for j := int32(0); j < mazeHeight; j++ {
			if i % 2 == 0 || j % 2 == 0 || i == 0 || j == 0 ||i == mazeWidth - 1 || j == mazeHeight - 1 {
				maze[i][j] = true
			} else {
				maze[i][j] = false
			}
		}
	}
	
	for i := 0; i < mazeHeight; i++ {
		for j := 0; j < mazeWidth; j++ {
			if maze[i][j] {
				tempMaze[i][j] = 2
			} else {
				tempMaze[i][j] = 0
			}
			
		}
	}
	
	
	rand.Seed(time.Now().UnixNano())
	r  := rand.Intn(mazeWidth / 2 + 1)
	r = r * 2 + 1
	
	boxX := int32(int32(r) * size)
	boxY := int32(0)
	tempMaze[r][0] = 1
	tempMaze = generateMaze(1, 1, tempMaze)
	
	r  = rand.Intn(mazeWidth / 2 + 1)
	r = r * 2 + 1
	tempMaze[r][mazeHeight - 1] = 1
	
	for i := 0; i < mazeHeight; i++ {
		for j := 0; j < mazeWidth; j++ {
			if tempMaze[i][j] == 2{
				maze[i][j] = true
			} else {
				maze[i][j] = false
			}
			
		}
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	temp := sdl.Rect{boxX, boxY, boxX + size, boxY + size}

	for i := int32(0); i < mazeHeight; i++ {
		for j := int32(0); j < mazeWidth; j++ {
			if maze[i][j] {
				temp = sdl.Rect{i * size, j * size, size, size}
				surface.FillRect(&temp, 0x3878e0)
			}
		}
	}

	rect := sdl.Rect{boxX, boxY, size, size}
	surface.FillRect(&rect, 0x579b29)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
				case *sdl.QuitEvent:
					println("Quit")
					running = false
					break
				case *sdl.KeyboardEvent:
					if t.State == 1 {
						if t.Keysym.Sym == 's'{
							if boxY < height - size {
								if !maze[boxX/ size][boxY / size + 1]{
									boxY = boxY + size
								}
							} else {
								fmt.Println("You won?")
								running = false
								break
							}
							
						}
						if t.Keysym.Sym == 'w' {
							if boxY > 0 {
								if !maze[boxX / size][boxY / size - 1]{
									boxY = boxY - size
								}
							} else {
								fmt.Println("You gave up. Smart.")
								running = false
								break
							}
						}
						if t.Keysym.Sym == 'd' && !maze[boxX / size + 1][boxY / size]{
							boxX = boxX + size
						}
						if t.Keysym.Sym == 'a' && !maze[boxX / size - 1][boxY / size]{
							boxX = boxX - size
						}
					}
					rect = sdl.Rect{boxX, boxY, size, size}
					surface.FillRect(nil, 0)
					for i := int32(0); i < mazeHeight; i++ {
						for j := int32(0); j < mazeWidth; j++ {
							if maze[i][j] {
								temp = sdl.Rect{i * size, j * size, size, size}
								surface.FillRect(&temp, 0x3878e0)
							}
						}
					}
					
					surface.FillRect(&rect, 0x579b29)
					window.UpdateSurface()
			}
		}
	}
}

func generateMaze (x int, y int, maze[mazeWidth][mazeHeight] int) [mazeWidth][mazeHeight] int {
	if len(maze) == 0 {
        fmt.Println("you broke it")
    }
	
	directions := make([]int, 0)
	
	if y > 1 {
		directions = append(directions, 0)
	}
	
	if y < mazeHeight - 2 {
		directions = append(directions, 2)
	}
	
	if x > 1 {
		directions = append(directions, 3)
	}
	
	if x < mazeWidth - 2 {
		directions = append(directions, 1)
	}

	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for i := 0; i < len(directions); i++ {
		switch (directions[i]) {
			//go up
			case 0: 
					if maze[x][y - 2] == 0{
						maze[x][y - 1] = 1
						maze[x][y - 2] = 1
						maze = generateMaze(x, y - 2, maze)
					}
			//go right
			case 1:
					if maze[x + 2][y] == 0{
						maze[x + 1][y] = 1
						maze[x + 2][y] = 1
						maze = generateMaze(x + 2, y, maze)
					}
			//go down
			case 2:
					if maze[x][y + 2] == 0{
						maze[x][y + 1] = 1
						maze[x][y + 2] = 1
						maze = generateMaze(x, y + 2, maze)
					}
			//go left
			case 3:
					if maze[x - 2][y] == 0{
						maze[x - 1][y] = 1
						maze[x - 2][y] = 1
						maze = generateMaze(x - 2, y, maze)
					}
		}
	}
	
	return maze 
}
