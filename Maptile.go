package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/x/fyne/widget"
	"github.com/fyne-io/terminal"
)

func main() {
	a := app.New()
	w := a.NewWindow("Terminal")
	m := widget.NewMap()
	t := terminal.New()
  	s := terminal.New()

	go func() {
		_ = t.RunLocalShell()
		a.Quit()
	}()
	go func() {
		_ = s.RunLocalShell()
		a.Quit()
	}()
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(container.NewGridWithColumns(
		3,
		t,
		m,
    s,
	))
	w.ShowAndRun()
}
