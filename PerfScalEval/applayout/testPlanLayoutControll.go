package applayout

import (
	"fmt"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"perfscaleval/confmodel"
)

func TestPlanControll() fyne.CanvasObject {
	// Zawartość dla sekcji Test Plan
	testCases := []string{"Test 1", "Test 2", "Test 3"}
	checkboxes := make([]fyne.CanvasObject, len(testCases))

	for i, tc := range testCases {
		checkboxes[i] = widget.NewCheck(tc, nil)
	}

	startTestBtn := widget.NewButton("Rozpocznij test", nil)
	progress := widget.NewProgressBar()

	container := container.NewVBox(
		container.NewVBox(checkboxes...),
		layout.NewSpacer(),
		progress,
		startTestBtn,
	)

	return container

}

func checkExistance() fyne.CanvasObject {

	if !confmodel.ProjectExist {
		return TestPlanCreate()
	} else if confmodel.ProjectExist {
		return TestPlanEdit()
	} else {
		return container.NewVBox(nil, nil, nil, nil)
	}

}

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
			confmodel.ProjectExist = true
			confmodel.CurrentPlan = &plan
			SwitchView(CreateTestPlanContent())
		})
		planButtons.Add(btn)
	}

	// Przycisk "Dodaj nowy"
	addBtn := widget.NewButton("+ Dodaj nowy Plan", func() {
		// Tutaj obsługa tworzenia nowego planu
		fmt.Println("Dodawanie nowego planu")
		confmodel.ProjectExist = false
		SwitchView(CreateTestPlanContent())

	})

	// Kontener końcowy - przyciski planów + przycisk dodawania
	return container.NewHScroll( // Dodaj scroll jeśli przycisków jest za dużo
		container.NewHBox(
			planButtons,
			addBtn,
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
				confmodel.ProjectExist = false
				SwitchView(CreateTestPlanContent())

			})
			return container.NewHBox(
				addBtn,
				widget.NewLabel("Brak planów"),
			)
		}

		// Dodaj przyciski dla każdego planu
		for _, group := range groups {
			group := group // ważne dla zamknięcia funkcji
			btn := widget.NewButton(group.Name, func() {
				// Tutaj obsługa kliknięcia w plan
				fmt.Println("Wybrano plan:", group.Name)
				confmodel.CurrentGroup = &group
				SwitchView(CreateTestPlanContent())
			})
			groupButtons.Add(btn)
		}

		// Przycisk "Dodaj nowy"
		addBtn := widget.NewButton("+ Dodaj nową Grupę", func() {
			// Tutaj obsługa tworzenia nowego planu
			fmt.Println("Dodawanie nowej grupy")
			confmodel.ProjectExist = false
			SwitchView(CreateTestPlanContent())

		})

		bar := container.NewHScroll(
			container.NewHBox(
				groupButtons,
				addBtn,
			),
		)
		return bar
	}

	return container.NewCenter(
		widget.NewLabel("Brak grup"),
	)

	// Kontener końcowy - przyciski planów + przycisk dodawania

}

func StepsBar() fyne.CanvasObject {
	// Kontener na przyciski planów
	stepButtons := container.NewHBox()
	if confmodel.CurrentPlan != nil {
		if confmodel.CurrentGroup != nil {
			// Pobierz listę planów z confmodel
			steps := confmodel.CurrentGroup.Steps

			// Jeśli nie ma planów, wyświetl napis
			if len(steps) == 0 {
				addBtn := widget.NewButton("+ Dodaj nowy Krok", func() {
					// Tutaj obsługa tworzenia nowego planu
					fmt.Println("Dodawanie nowego kroku")
					SwitchView(CreateTestPlanContent())

				})
				return container.NewHBox(
					addBtn,
					widget.NewLabel("Brak planów"),
				)
			}

			// Dodaj przyciski dla każdego planu
			for _, step := range steps {
				step := step // ważne dla zamknięcia funkcji
				btn := widget.NewButton(step.Name, func() {
					// Tutaj obsługa kliknięcia w plan
					fmt.Println("Wybrano krok:", step.Name)
					confmodel.CurrentStep = &step
					SwitchView(CreateTestPlanContent())
				})
				stepButtons.Add(btn)
			}

			// Przycisk "Dodaj nowy"
			addBtn := widget.NewButton("+ Dodaj nowy Krok", func() {
				// Tutaj obsługa tworzenia nowego planu
				fmt.Println("Dodawanie nowego kroku")
				SwitchView(CreateTestPlanContent())

			})

			bar := container.NewVScroll(
				container.NewVBox(
					stepButtons,
					addBtn,
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

func CreateTestPlanContent() fyne.CanvasObject {
	side := StepsBar()

	changer := checkExistance()

	MainContent := container.NewBorder(GroupBar(), nil, side, nil, changer)

	PlanContent := container.NewVBox(PlanBar(), MainContent)

	return PlanContent
}
