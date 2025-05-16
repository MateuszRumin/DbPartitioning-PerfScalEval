package applayout

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Globalna zmienna do przechowywania referencji

func CreateTopMenu(content *fyne.Container) fyne.CanvasObject {
	connectionsBtn := widget.NewButton("Połączenia", func() {
		content.Objects = []fyne.CanvasObject{CreateConnectionsContent()}
		content.Refresh()
	})

	testPlanBtn := widget.NewButton("Test Plan", func() {
		content.Objects = []fyne.CanvasObject{CreateTestPlanContent()}
		content.Refresh()
	})

	resultsBtn := widget.NewButton("Wyniki", func() {
		content.Objects = []fyne.CanvasObject{CreateResultsContent()}
		content.Refresh()
	})

	return container.NewGridWithColumns(3, connectionsBtn, testPlanBtn, resultsBtn)
}
