package applayout

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"perfscaleval/confmodel"
)

func stepControler() fyne.CanvasObject {
	if confmodel.CreateMode {
		return StepCreate()
	}

	return stepEdit()

}
func groupControler() fyne.CanvasObject {
	if confmodel.CreateMode {
		return GroupCreate()
	}

	return GroupEdit()
}

func planControler() fyne.CanvasObject {
	if confmodel.CreateMode {
		return TestPlanCreate()
	}

	return TestPlanEdit()
}

func planViewMode() fyne.CanvasObject {

	if confmodel.ChooseLayOut == 0 {
		return planControler()
	} else if confmodel.ChooseLayOut == 1 {
		return groupControler()
	} else if confmodel.ChooseLayOut == 2 {
		return stepControler()
	} else {
		return widget.NewLabel("Błąd")
	}

}

func CreateTestPlanContent() fyne.CanvasObject {

	side := StepsBar()

	changer := planViewMode()

	MainContent := container.NewBorder(GroupBar(), nil, side, nil, container.NewStack(changer))
	MainContent.Resize(fyne.NewSize(MainContent.MinSize().Width, 2000))

	PlanContent := container.NewVBox(PlanBar(), MainContent)

	return PlanContent
}
