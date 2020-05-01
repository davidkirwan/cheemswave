package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	resources "github.com/davidkirwan/parallax_scrolling/internal/pkg/resources"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func init() {
	//
}

func removeTree(s []*pixel.Sprite, index int) []*pixel.Sprite {
	return append(s[:index], s[index+1:]...)
}

func removeMatrices(s []pixel.Matrix, index int) []pixel.Matrix {
	return append(s[:index], s[index+1:]...)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spritesheet, err := resources.LoadPNGPicture("assets/images/trees.png")
	if err != nil {
		panic(err)
	}

	var treesFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	var (
		camPos       = pixel.ZV
		camSpeed     = 500.0
		camZoom      = 1.0
		camZoomSpeed = 1.2
		// camXStart    = pixel.ZV.X
		// camYStart    = pixel.ZV.Y
		trees    []*pixel.Sprite
		matrices []pixel.Matrix
		counter  = 0
	)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])
			trees = append(trees, tree)
			mouse := cam.Unproject(win.MousePosition())
			matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
		}

		if counter == 250 {
			counter = 0
		} else {
			counter++
			//camPos.X -= camSpeed * dt

			if counter%50 == 0 {
				fmt.Printf("Counter: %d, X: %f, Y: %f\n", counter, camPos.X, camPos.Y)
				x := camPos.X - 1000.0
				y := camPos.Y + rand.Float64()*300
				tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])

				treeVec := &pixel.Vec{X: x, Y: y}
				tree.Frame().Moved(*treeVec)
				trees = append(trees, tree)
				matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(*treeVec))
			}
		}

		if win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA) {
			camPos.X -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyD) {
			camPos.X += camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) {
			camPos.Y -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		win.Clear(colornames.Forestgreen)

		markedForRemoval := []int{}

		for i, tree := range trees {
			tree.Draw(win, matrices[i])
			tb := tree.Picture().Bounds().Center()
			if resources.CalculateDistance(camPos.X, camPos.Y, tb.X, tb.Y) {
				markedForRemoval = append(markedForRemoval, i)
			}
		}

		for i := len(markedForRemoval) - 1; i >= 0; i-- {
			fmt.Printf("size matrix: %d, size trees: %d\n", len(matrices), len(trees))
			matrices = removeMatrices(matrices, markedForRemoval[i])
			trees = removeTree(trees, markedForRemoval[i])
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
