package main

import (
	"perfscaleval/applayout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var (
	mainWindow  fyne.Window
	mainContent *fyne.Container
)

func main() {
	a := app.New()
	mainWindow = a.NewWindow("Moja Aplikacja Testowa")
	mainWindow.Resize(fyne.NewSize(1024, 768))

	initMainContent()
	mainWindow.ShowAndRun()
}

func initMainContent() {
	mainContent = container.NewStack()

	// Inicjalizacja manager widoków
	applayout.InitViewManager(mainContent)

	topMenu := applayout.LayOut() // Teraz bez callbacka
	content := container.NewBorder(topMenu, nil, nil, nil, mainContent)

	// Ładowanie domyślnego widoku
	applayout.SwitchView(applayout.CreateConnectionsContent())

	mainWindow.SetContent(content)
}
