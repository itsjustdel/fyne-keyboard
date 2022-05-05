package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// var format beep.Format
// var streamer beep.StreamSeekCloser

func main() {

	a := app.New()
	w := a.NewWindow("Lorenzo Harmonium")

	t := widget.NewLabel("Lorenzo Harmonium")

	//samples
	f, err := os.Open("ALittleMore.mp3")
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

	keyTotal := 54
	var keys = []fyne.CanvasObject{}

	var width float32 = 50
	for i := 0; i < keyTotal; i++ {
		//button with empty function body
		k := widget.NewButton("", func() {})

		//need to set this after declaration so it prints correct number from label
		//if we print "i" in new button declaration, it used i = 54 (or whatever the last index was)
		k.OnTapped = func() {
			//need to find instance of key
			for j := 0; j < keyTotal; j++ {

				if k == keys[j] {
					fmt.Println(fmt.Sprint(j))
					//play sound
					//desktop.Canvas.OnKeyDown()
					speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
					speaker.Play(streamer)
				}
			}
		}

		k.Move(fyne.NewPos(float32(i)*width, 50))
		k.Resize(fyne.NewSize(width, 300))

		keys = append(keys, fyne.CanvasObject(k))

	}

	//bottomC := container.NewHBox(keys...)

	noLayout := container.NewWithoutLayout(keys...)
	c := container.NewVBox(t, noLayout)

	w.SetContent(c)
	w.Resize(fyne.NewSize(1600, 500))
	w.ShowAndRun()

}
