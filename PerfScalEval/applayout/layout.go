package applayout

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Globalna zmienna do przechowywania referencji

func LayOut() fyne.CanvasObject {
	savePlan := widget.NewButton("Zapisz", func() {})
	loadPlan := widget.NewButton("Wczytaj", func() {})

	saveings := container.NewVBox(savePlan, loadPlan)

	connectionsBtn := widget.NewButton("Połączenia", func() {
		SwitchView(CreateConnectionsContent())
	})

	testPlanBtn := widget.NewButton("Test Plan", func() {
		SwitchView(CreateTestPlanContent())
	})
	resultsBtn := widget.NewButton("Wyniki", func() {
		SwitchView(CreateResultsContent())
	})

	return container.NewGridWithColumns(4, connectionsBtn, testPlanBtn, resultsBtn, saveings)
}
