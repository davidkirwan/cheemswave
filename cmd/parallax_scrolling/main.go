/*
    Cheemswave - Side Scroller Game
    Copyright (C) 2021  David Kirwan

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"math"
	"os"
	"time"
	"strconv"

	resources "github.com/davidkirwan/parallax_scrolling/internal/pkg/resources"
//	game "github.com/davidkirwan/parallax_scrolling/internal/pkg/game"

	"github.com/faiface/beep"
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
	f, err := os.Open("assets/audio/Night_Lights_Original_Mix.mp3")
	if err != nil {
		log.Fatal(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	done := make(chan bool)
	streamer, format, err := mp3.Decode(f)
	loop := beep.Loop(3, streamer)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))

	cfg := pixelgl.WindowConfig{
		Title:  "Cheemswave - Night Lights (Original Mix) - created with Pixel",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spritesheet, err := resources.LoadPNGPicture("assets/images/virus.png")
	if err != nil {
		panic(err)
	}

	c, err := resources.LoadPNGPicture("assets/images/cheems_forward.png")
	if err != nil {
		panic(err)
	}
	cheems := pixel.NewSprite(c, c.Bounds())

	c2, err := resources.LoadPNGPicture("assets/images/cheems_backward.png")
	if err != nil {
		panic(err)
	}
	cheemsBack := pixel.NewSprite(c2, c2.Bounds())

	b, err := resources.LoadPNGPicture("assets/images/bork.png")
	if err != nil {
		panic(err)
	}
	bork := pixel.NewSprite(b, b.Bounds())

	var viriiFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			viriiFrames = append(viriiFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	var (
		camPos    = pixel.ZV
		cheemsVec = &pixel.Vec{X: camPos.X - 400, Y: camPos.Y - 200}
		borkVec   = &pixel.Vec{X: cheemsVec.X - 1000, Y: cheemsVec.Y}
		camSpeed  = 400.0
		camZoom   = 1.0
		camZoomSpeed = 1.2
		//camXStart    = pixel.ZV.X
		//camYStart    = pixel.ZV.Y
		virii          []*pixel.Sprite
		matrices       []pixel.Matrix
		counter        = 0
		viriiCounter    = 0
		cheemsBackward = false
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
		y := camPos.Y + rand.Float64() * 768 - 300

		fmt.Printf("Counter: %d, X: %f, Y: %f\n", counter, camPos.X, len(virii))
		if counter == 400 {
			counter = 0
			viriiCounter = 0
		}
		if counter%15 == 0 {
			if len(virii) < 50 {
				virus := pixel.NewSprite(spritesheet, viriiFrames[rand.Intn(len(viriiFrames))])
				viriiVec := &pixel.Vec{X: x, Y: y}
				virii = append(virii, virus)
				matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(*viriiVec))
			} else {
				viriiVec := &pixel.Vec{X: x, Y: y}
				matrices[viriiCounter] = pixel.IM.Scaled(pixel.ZV, 4).Moved(*viriiVec)
				viriiCounter++
			}
		}
		counter++

		if win.Pressed(pixelgl.KeyP) {
			if camSpeed == 0 {
				camSpeed = 400
			} else {
				camSpeed = 0
			}
		}

		if win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA) {
		        if debug {
				camPos.X -= camSpeed * dt
			}
			cheemsBackward = true
			cheemsVec = &pixel.Vec{X: cheemsVec.X - (camSpeed*0.2)*dt, Y: cheemsVec.Y}
		}
		if win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyD) {
			if debug {
			  camPos.X += camSpeed * dt
			}
			cheemsBackward = false
			cheemsVec = &pixel.Vec{X: cheemsVec.X + (camSpeed*2)*dt, Y: cheemsVec.Y}
		}
		if win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) {
			if debug {
			  camPos.Y -= camSpeed * dt
			}
			cheemsVec = &pixel.Vec{X: cheemsVec.X, Y: cheemsVec.Y - camSpeed*dt}
		}
		if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) {
			if debug {
			  camPos.Y += camSpeed * dt
			}
			cheemsVec = &pixel.Vec{X: cheemsVec.X, Y: cheemsVec.Y + camSpeed*dt}
		}
		if debug {
		  camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)
		}

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

		for i, virus := range virii {
			virus.Draw(win, matrices[i])
		}

		if cheemsBackward {
			cheemsBack.Draw(win, pixel.IM.Scaled(pixel.ZV, 0.15).Moved(*cheemsVec))
		} else {
			cheems.Draw(win, pixel.IM.Scaled(pixel.ZV, 0.15).Moved(*cheemsVec))
		}

		bork.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(*borkVec))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
