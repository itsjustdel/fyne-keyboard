package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func uIKeys() *fyne.Container {

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

					//trigger play sound
				}
			}
		}

		k.Move(fyne.NewPos(float32(i)*width, 50))
		k.Resize(fyne.NewSize(width, 300))

		keys = append(keys, fyne.CanvasObject(k))
	}

	noLayout := container.NewWithoutLayout(keys...)

	return noLayout
}
