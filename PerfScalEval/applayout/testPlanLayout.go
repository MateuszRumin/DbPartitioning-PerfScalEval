package applayout

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateTestPlanContent() fyne.CanvasObject {
	// Zawartość dla sekcji Test Plan
	testCases := []string{"Test 1", "Test 2", "Test 3"}
	checkboxes := make([]fyne.CanvasObject, len(testCases))

	for i, tc := range testCases {
		checkboxes[i] = widget.NewCheck(tc, nil)
	}

	startTestBtn := widget.NewButton("Rozpocznij test", nil)
	progress := widget.NewProgressBar()

	return container.NewVBox(
		container.NewVBox(checkboxes...),
		layout.NewSpacer(),
		progress,
		startTestBtn,
	)
}
