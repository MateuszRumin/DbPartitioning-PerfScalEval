package applayout

import (
	"fmt"
	"perfscaleval/confmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func StepCreate() fyne.CanvasObject {

	nameEntry := widget.NewEntry()
	nameContainer := container.NewVBox(widget.NewLabel("Nazwa"), nameEntry)

	queryEntry := widget.NewEntry()
	queryContainer := container.NewVBox(widget.NewLabel("Zapytanie"), queryEntry)

	saveBtn := widget.NewButton("Dodaj plan", func() {
		if nameEntry.Text == "" || queryEntry.Text == "" {
			dialog.ShowInformation("Błąd", "Nazwa planu nie może być pusta!", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		newStep := confmodel.TestStep{
			Name:  nameEntry.Text,
			Query: queryEntry.Text,
		}

		confmodel.CurrentGroup.Steps = append(confmodel.CurrentGroup.Steps, newStep)
		confmodel.CurrentStep = &confmodel.CurrentGroup.Steps[len(confmodel.Plan)-1]
		confmodel.ChooseLayOut = 2
		confmodel.CreateMode = false

		// Logowanie dla debugowania (można usunąć w produkcji)
		fmt.Println("Zmodtfikowano plan")
		fmt.Printf("Aktualne plany: %+v\n", confmodel.Plan)
		SwitchView(CreateTestPlanContent())
	})

	mainContainer := container.NewVBox(nameContainer, queryContainer, saveBtn)

	return mainContainer

}

func stepEdit() fyne.CanvasObject {
	confmodel.CurrentStep = &confmodel.CurrentGroup.Steps[0]
	nameEntry := widget.NewEntry()
	nameEntry.SetText(confmodel.CurrentStep.Name)
	nameContainer := container.NewVBox(widget.NewLabel("Nazwa"), nameEntry)

	queryEntry := widget.NewEntry()
	queryEntry.SetText(confmodel.CurrentStep.Query)
	queryContainer := container.NewVBox(widget.NewLabel("Zapytanie"), queryEntry)

	saveBtn := widget.NewButton("Edytuj plan", func() {
		if nameEntry.Text == "" || queryEntry.Text == "" {
			dialog.ShowInformation("Błąd", "Nazwa planu nie może być pusta!", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		newStep := confmodel.TestStep{
			Name:  nameEntry.Text,
			Query: queryEntry.Text,
		}

		confmodel.CurrentGroup.Steps = append(confmodel.CurrentGroup.Steps, newStep)
		confmodel.CurrentStep = &confmodel.CurrentGroup.Steps[len(confmodel.Plan)-1]
		confmodel.ChooseLayOut = 2
		confmodel.CreateMode = false

		// Logowanie dla debugowania (można usunąć w produkcji)
		fmt.Println("Zmodtfikowano plan")
		fmt.Printf("Aktualne plany: %+v\n", confmodel.Plan)
		SwitchView(CreateTestPlanContent())
	})

	mainContainer := container.NewVBox(nameContainer, queryContainer, saveBtn)

	return mainContainer

}
