package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"perfscaleval/applayout"
)

var mainContent *fyne.Container

func main() {
	a := app.New()
	w := a.NewWindow("Moja Aplikacja Testowa")
	w.Resize(fyne.NewSize(1024, 768))

	// Inicjalizacja głównej zawartości
	mainContent = container.NewStack()

	// Utwórz i podłącz menu
	topMenu := applayout.CreateTopMenu(mainContent)

	// Kompozycja layoutu
	content := container.NewBorder(topMenu, nil, nil, nil, mainContent)

	// Domyślna zawartość
	mainContent.Objects = []fyne.CanvasObject{applayout.CreateConnectionsContent()}

	w.SetContent(content)
	w.ShowAndRun()
}
