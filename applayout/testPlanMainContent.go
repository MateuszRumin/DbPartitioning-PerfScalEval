package applayout

import (
	"fmt"
	"perfscaleval/confmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func TestPlanCreate() fyne.CanvasObject {
	// Pole wprowadzania dla nazwy
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Wprowadź nazwę...")

	// Pole wprowadzania dla komentarza
	commentEntry := widget.NewMultiLineEntry()
	commentEntry.SetPlaceHolder("Wprowadź komentarz...")

	// Przycisk do zapisu
	saveBtn := widget.NewButton("Stwórz plan", func() {
		if nameEntry.Text == "" {
			dialog.ShowInformation("Błąd", "Nazwa planu nie może być pusta!", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		// Tworzymy nowy TestPlan z wprowadzonych danych
		newPlan := confmodel.TestPlan{
			PlanName: nameEntry.Text,
			Comments: commentEntry.Text,
			Group:    []confmodel.TestGroup{}, // Początkowo pusta lista grup
		}
		confmodel.CreateMode = false

		// Dodajemy do globalnej zmiennej Plan
		confmodel.Plan = append(confmodel.Plan, newPlan)
		confmodel.CurrentPlan = &confmodel.Plan[len(confmodel.Plan)-1]

		// Logowanie dla debugowania (można usunąć w produkcji)
		fmt.Printf("Zapisano nowy plan: %+v\n", newPlan)
		fmt.Printf("Aktualne plany: %+v\n", confmodel.Plan)
		SwitchView(CreateTestPlanContent())
	})

	// Kontener z elementami ułożonymi pionowo
	container := container.NewVBox(
		widget.NewLabel("Nazwa:"),
		nameEntry,
		widget.NewLabel("Komentarz:"),
		commentEntry,
		saveBtn,
	)

	return container
}

func TestPlanEdit() fyne.CanvasObject {

	// Pole wprowadzania dla nazwy
	nameEntry := widget.NewEntry()
	nameEntry.SetText(confmodel.CurrentPlan.PlanName)

	// Pole wprowadzania dla komentarza
	commentEntry := widget.NewMultiLineEntry()
	commentEntry.SetText(confmodel.CurrentPlan.Comments)

	// Przycisk do zapisu
	saveBtn := widget.NewButton("Zmodyfikuj plan", func() {
		if nameEntry.Text == "" {
			dialog.ShowInformation("Błąd", "Nazwa planu nie może być pusta!", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		confmodel.CurrentPlan.PlanName = nameEntry.Text
		confmodel.CurrentPlan.Comments = commentEntry.Text

		// Logowanie dla debugowania (można usunąć w produkcji)
		fmt.Println("Zmodtfikowano plan")
		fmt.Printf("Aktualne plany: %+v\n", confmodel.Plan)
		SwitchView(CreateTestPlanContent())
	})
	delateBtn := widget.NewButton("Usuń plan", func() {
		for i, plan := range confmodel.Plan {
			if plan.PlanName == confmodel.CurrentPlan.PlanName {
				confmodel.Plan = append(confmodel.Plan[:i], confmodel.Plan[i+1:]...)
				break
			}
		}

		if len(confmodel.Plan) > 0 {
			confmodel.CurrentPlan = &confmodel.Plan[0]
		} else {
			confmodel.CurrentPlan = nil
			confmodel.CurrentGroup = nil
			confmodel.CurrentStep = nil
			confmodel.ChooseLayOut = 0
			confmodel.CreateMode = true
		}
		SwitchView(CreateTestPlanContent())
	})
	buttons := container.NewHBox(saveBtn, delateBtn)

	// Kontener z elementami ułożonymi pionowo
	container := container.NewVBox(
		widget.NewLabel("Nazwa:"),
		nameEntry,
		widget.NewLabel("Komentarz:"),
		commentEntry,
		buttons,
	)

	return container
}
