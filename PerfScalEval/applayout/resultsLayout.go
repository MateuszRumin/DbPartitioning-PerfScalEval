package applayout

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateResultsContent() fyne.CanvasObject {
	// Zawartość dla sekcji Wyniki
	data := [][]string{
		{"Test 1", "SUKCES", "2023-10-01"},
		{"Test 2", "BŁĄD", "2023-10-01"},
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 3 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.SetText(data[id.Row][id.Col])
		},
	)

	refreshBtn := widget.NewButton("Odśwież", nil)
	clearBtn := widget.NewButton("Wyczyść", nil)

	return container.NewBorder(
		container.NewHBox(refreshBtn, clearBtn),
		nil,
		nil,
		nil,
		table,
	)
}
