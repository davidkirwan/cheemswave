package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	resources "github.com/davidkirwan/parallax_scrolling/internal/pkg/resources"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func removeTree(s []*pixel.Sprite, index int) []*pixel.Sprite {
	return append(s[:index], s[index+1:]...)
}

func removeMatrices(s []pixel.Matrix, index int) []pixel.Matrix {
	return append(s[:index], s[index+1:]...)
}

func run() {
	f, err := os.Open("assets/sound/Night_Lights_Original_Mix.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	cfg := pixelgl.WindowConfig{
		Title:  "Cheemswave - Night Lights (Original Mix) - created with Pixel",
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

	c, err := resources.LoadPNGPicture("assets/images/cheems.png")
	if err != nil {
		panic(err)
	}
	cheems := pixel.NewSprite(c, c.Bounds())

	b, err := resources.LoadPNGPicture("assets/images/bork.png")
	if err != nil {
		panic(err)
	}
	bork := pixel.NewSprite(b, b.Bounds())

	var treesFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	var (
		camPos    = pixel.ZV
		cheemsVec = &pixel.Vec{X: camPos.X - 400, Y: camPos.Y - 200}
		borkVec   = &pixel.Vec{X: cheemsVec.X - 1000, Y: cheemsVec.Y}
		camSpeed  = 400.0
		camZoom   = 1.0
		//camZoomSpeed = 1.2
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
			borkVec = &pixel.Vec{X: cheemsVec.X + 30, Y: cheemsVec.Y + 50}
		}

		camPos.X += camSpeed * dt
		x := camPos.X + 1000.0
		y := camPos.Y + rand.Float64()*600 - 300

		if counter == 250 {
			counter = 0
		} else {
			counter++

			if counter%20 == 0 {
				fmt.Printf("Counter: %d, X: %f, Y: %f\n", counter, camPos.X, camPos.Y)
				tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])

				treeVec := &pixel.Vec{X: x, Y: y}
				tree.Frame().Moved(*treeVec)
				trees = append(trees, tree)
				matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(*treeVec))
			}
		}

		if win.Pressed(pixelgl.KeyP) {
			if camSpeed == 0 {
				camSpeed = 400
			} else {
				camSpeed = 0
			}
		}

		if win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA) {
			//camPos.X -= camSpeed * dt
			cheemsVec = &pixel.Vec{X: cheemsVec.X - (camSpeed*0.2)*dt, Y: cheemsVec.Y}
		}
		if win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyD) {
			//camPos.X += camSpeed * dt
			cheemsVec = &pixel.Vec{X: cheemsVec.X + (camSpeed*2)*dt, Y: cheemsVec.Y}
		}
		if win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) {
			//camPos.Y -= camSpeed * dt
			cheemsVec = &pixel.Vec{X: cheemsVec.X, Y: cheemsVec.Y - camSpeed*dt}
		}
		if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) {
			//camPos.Y += camSpeed * dt
			cheemsVec = &pixel.Vec{X: cheemsVec.X, Y: cheemsVec.Y + camSpeed*dt}
		}
		//camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		if cheemsVec.X < camPos.X-450 {
			cheemsVec = &pixel.Vec{X: camPos.X - 450, Y: cheemsVec.Y}
		}

		if cheemsVec.Y > camPos.Y+350 {
			cheemsVec = &pixel.Vec{X: cheemsVec.X, Y: camPos.Y + 350}
		}

		if cheemsVec.Y < camPos.Y-350 {
			cheemsVec = &pixel.Vec{X: cheemsVec.X, Y: camPos.Y - 350}
		}

		win.Clear(colornames.Azure)

		for i, tree := range trees {
			tree.Draw(win, matrices[i])
		}

		cheems.Draw(win, pixel.IM.Scaled(pixel.ZV, 0.05).Moved(*cheemsVec))
		bork.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(*borkVec))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
