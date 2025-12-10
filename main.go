package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	prometheusGoAppMetrics()

	aplication := app.New()
	window := aplication.NewWindow("Update Time")
	window.Resize(fyne.NewSize(1024, 500))

	buttonStartMesure := widget.NewButton("Start", func() {
		startBackgroundScraper()
	})
	buttonStartMesure.Resize(fyne.NewSize(220, 50))
	buttonStopMesure := widget.NewButton("Stop", func() {

		stopBackgroundScraper()

	})
	buttonStopMesure.Resize(fyne.NewSize(220, 50))

	MesureButtonContainer := container.NewHBox(buttonStartMesure, buttonStopMesure)

	MessureActionContainer := container.NewCenter(MesureButtonContainer)

	MainContainer := container.NewBorder(container.NewCenter(widget.NewLabel("Pomiarowanie")), MessureActionContainer, nil, nil, nil)

	window.SetContent(MainContainer)
	window.ShowAndRun()

}
