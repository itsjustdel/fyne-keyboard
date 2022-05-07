package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Lorenzo Harmonium")

	l := widget.NewLabel("Lorenzo Harmonium")

	setKeys(w)

	c := container.NewVBox(l, uIKeys())

	w.SetContent(c)

	w.Resize(fyne.NewSize(1600, 500))
	w.ShowAndRun()

}
