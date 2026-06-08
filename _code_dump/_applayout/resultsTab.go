package applayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"perfscaleval/tests"
)

func CreateResultsContent() fyne.CanvasObject {

	planBar := PlanBarResults()
	groupBar := GroupBarResults()

	btn_start := widget.NewButton("Start", func() {
		tests.CheckEnabledConnection()
	})

	containerUp := container.NewVBox(planBar, groupBar)

	container := container.NewBorder(containerUp, btn_start, nil, nil, nil)

	return container
}
