package applayout

import (
	"fmt"
	"perfscaleval/confmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PlanBar() fyne.CanvasObject {
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
			confmodel.ChooseLayOut = 0
			confmodel.CreateMode = false
			confmodel.CurrentPlan = &plan
			SwitchView(CreateTestPlanContent())
		})
		planButtons.Add(btn)
	}

	// Przycisk "Dodaj nowy"
	addBtn := widget.NewButton("+ Dodaj nowy Plan", func() {
		// Tutaj obsługa tworzenia nowego planu
		fmt.Println("Dodawanie nowego planu")
		confmodel.CreateMode = true
		confmodel.ChooseLayOut = 0
		SwitchView(CreateTestPlanContent())

	})

	// Kontener końcowy - przyciski planów + przycisk dodawania
	return container.NewBorder(
		nil, nil,
		nil, addBtn,
		container.NewHScroll( // Dodaj scroll jeśli przycisków jest za dużo
			container.NewHBox(
				planButtons,
			),
		),
	)
}

func GroupBar() fyne.CanvasObject {
	// Kontener na przyciski planów
	groupButtons := container.NewHBox()
	if confmodel.CurrentPlan != nil {
		// Pobierz listę planów z confmodel
		groups := confmodel.CurrentPlan.Group

		// Jeśli nie ma planów, wyświetl napis
		if len(groups) == 0 {
			addBtn := widget.NewButton("+ Dodaj nową Grupę", func() {
				// Tutaj obsługa tworzenia nowego planu
				fmt.Println("Dodawanie nowej grupy")
				confmodel.CreateMode = true
				confmodel.ChooseLayOut = 1
				SwitchView(CreateTestPlanContent())

			})
			return container.NewBorder(nil, nil, nil, addBtn, nil)
		}

		// Dodaj przyciski dla każdego planu
		for _, group := range groups {
			group := group // ważne dla zamknięcia funkcji
			btn := widget.NewButton(group.Name, func() {
				confmodel.CurrentGroup = &group
				confmodel.CreateMode = false
				confmodel.ChooseLayOut = 1
				SwitchView(CreateTestPlanContent())
			})
			groupButtons.Add(btn)
		}

		// Przycisk "Dodaj nowy"
		addBtn := widget.NewButton("+ Dodaj nową Grupę", func() {
			// Tutaj obsługa tworzenia nowego planu
			fmt.Println("Dodawanie nowej grupy")
			confmodel.CreateMode = true
			confmodel.ChooseLayOut = 1
			SwitchView(CreateTestPlanContent())

		})

		bar := container.NewBorder(nil, nil, nil, addBtn,
			container.NewHScroll(
				container.NewHBox(
					groupButtons,
				),
			))

		return bar
	}

	return container.NewCenter(
		widget.NewLabel("Brak grup"),
	)

	// Kontener końcowy - przyciski planów + przycisk dodawania

}

func StepsBar() fyne.CanvasObject {
	// Kontener na przyciski planów
	stepButtons := container.NewVBox()
	if confmodel.CurrentPlan != nil {
		if confmodel.CurrentGroup != nil {
			// Pobierz listę planów z confmodel
			steps := confmodel.CurrentGroup.Steps

			// Jeśli nie ma planów, wyświetl napis
			if len(steps) == 0 {
				addBtn := widget.NewButton("+ Dodaj nowy Krok", func() {
					// Tutaj obsługa tworzenia nowego planu
					fmt.Println("Dodawanie nowego kroku")
					confmodel.CreateMode = true
					confmodel.ChooseLayOut = 2
					SwitchView(CreateTestPlanContent())

				})
				return container.NewBorder(
					widget.NewLabel("Brak krokó"),
					addBtn,
					nil,
					nil,
					nil,
				)
			}

			// Dodaj przyciski dla każdego planu
			for _, step := range steps {
				step := step // ważne dla zamknięcia funkcji
				btn := widget.NewButton(step.Name, func() {
					// Tutaj obsługa kliknięcia w plan
					fmt.Println("Wybrano krok:", step.Name)
					confmodel.CreateMode = false
					confmodel.ChooseLayOut = 2
					confmodel.CurrentStep = &step
					SwitchView(CreateTestPlanContent())
				})
				stepButtons.Add(btn)
			}

			// Przycisk "Dodaj nowy"
			addBtn := widget.NewButton("+ Dodaj nowy Krok", func() {
				// Tutaj obsługa tworzenia nowego planu
				fmt.Println("Dodawanie nowego kroku")
				confmodel.CreateMode = true
				confmodel.ChooseLayOut = 2
				SwitchView(CreateTestPlanContent())

			})

			bar := container.NewBorder(
				widget.NewLabel("Kroki"),
				addBtn,
				nil,
				nil,
				container.NewVScroll(
					container.NewVBox(
						stepButtons,
					),
				),
			)
			return bar
		}

		return container.NewCenter(
			widget.NewLabel("Brak kroków"),
		)
	}

	return container.NewCenter(
		widget.NewLabel("Brak kroków"),
	)

	// Kontener końcowy - przyciski planów + przycisk dodawania

}
