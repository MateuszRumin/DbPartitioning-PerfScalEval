package applayout

import (
	"fmt"
	"perfscaleval/confmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PlanBarResults() fyne.CanvasObject {
	// Kontener na przyciski planów
	planButtons := container.NewHBox()

	// Pobierz listę planów z confmodel
	plans := confmodel.Plan

	// Jeśli nie ma planów, wyświetl napis
	if len(plans) == 0 {
		return container.NewCenter(
			widget.NewLabel("Brak planów"),
		)
	}

	// Dodaj przyciski dla każdego planu
	for _, plan := range plans {
		plan := plan // ważne dla zamknięcia funkcji
		btn := widget.NewButton(plan.PlanName, func() {
			// Tutaj obsługa kliknięcia w plan
			fmt.Println("Wybrano plan:", plan.PlanName)
		})
		planButtons.Add(btn)
	}

	// Kontener końcowy - przyciski planów + przycisk dodawania
	return container.NewHScroll( // Dodaj scroll jeśli przycisków jest za dużo
		container.NewHBox(
			planButtons,
		),
	)
}

func GroupBarResults() fyne.CanvasObject {
	// Kontener na przyciski planów
	groupButtons := container.NewHBox()
	if confmodel.CurrentPlan != nil {
		// Pobierz listę planów z confmodel
		groups := confmodel.CurrentPlan.Group

		// Jeśli nie ma planów, wyświetl napis
		if len(groups) == 0 {
			return widget.NewLabel("Brak Grup")
		}
		// Dodaj przyciski dla każdego planu
		for _, group := range groups {
			group := group // ważne dla zamknięcia funkcji
			btn := widget.NewButton(group.Name, func() {
			})
			groupButtons.Add(btn)
		}

		bar := container.NewHScroll(
			container.NewHBox(
				groupButtons),
		)
		return bar
	}

	return container.NewCenter(
		widget.NewLabel("Brak grup"),
	)

	// Kontener końcowy - przyciski planów + przycisk dodawania

}
