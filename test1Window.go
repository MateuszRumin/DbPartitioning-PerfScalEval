package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var test1Window fyne.Window

func test1PopUP(a fyne.App) {

	test1Window = a.NewWindow("Test 1")
	test1Window.Resize(fyne.NewSize(300, 300))

	saveBTN := widget.NewButton("Zakończ", func() {

		test1Window.Close()

	})

	buttons := container.NewHBox(saveBTN)

	connection := container.NewBorder(
		container.NewCenter(widget.NewLabel("Select Dataset")), // góra
		container.NewCenter(buttons),                           // dół
		nil,                                                    // lewo
		nil,                                                    // prawo
		nil,
	)

	test1Window.SetContent(connection)

	test1Window.Show()

}
