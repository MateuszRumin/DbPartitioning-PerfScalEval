package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	mainWindow  fyne.Window
	mainContent *fyne.Container
)

func main() {
	a := app.New()
	mainWindow = a.NewWindow("Perf-Scal-eval")
	mainWindow.Resize(fyne.NewSize(850, 500))

	initMainContent(a)

	mainWindow.Show()
	a.Run()
}

func initMainContent(a fyne.App) {
	mainContent = container.NewStack()

	topMenu := topMenu(a) // Teraz bez callbacka
	content := container.NewBorder(topMenu, nil, nil, nil, mainContent)

	// Ładowanie domyślnego widoku
	// applayout.SwitchView(applayout.CreateConnectionsContent())

	mainWindow.SetContent(content)
}

func topMenu(a fyne.App) fyne.CanvasObject {

	connectionsBtn := widget.NewButton("Połączenie", func() {
		setConnectionValues(a)

	})
	connectionsBtn.Resize(fyne.NewSize(220, 50))

	topMenu := container.NewHBox(
		connectionsBtn,
	)

	return container.NewCenter(topMenu)
}
