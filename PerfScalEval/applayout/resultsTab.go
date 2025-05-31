package applayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"perfscaleval/selectests"
)

func CreateResultsContent() fyne.CanvasObject {

	planBar := PlanBarResults()
	groupBar := GroupBarResults()

	btn_start := widget.NewButton("Start", func() {
		selectests.RunTest()
	})

	containerUp := container.NewVBox(planBar, groupBar)

	container := container.NewBorder(containerUp, btn_start, nil, nil, nil)

	return container
}
