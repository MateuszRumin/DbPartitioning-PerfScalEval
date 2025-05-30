package applayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StepCreate() fyne.CanvasObject {
	checkRamUp := widget.NewCheck("RampUp", func(checked bool) {
		if checked {
			println("Checkbox is checked")
		} else {
			println("Checkbox is unchecked")
		}
	})

	checkSameUserIter := widget.NewCheck("Same User Iter", func(checked bool) {
		if checked {
			println("Checkbox is checked")
		} else {
			println("Checkbox is unchecked")
		}
	})
	infiniteIter := widget.NewCheck("Infinite Loop", func(checked bool) {
		if checked {
			println("Checkbox is checked")
		} else {
			println("Checkbox is unchecked")
		}
	})

	checksContainer := container.NewHBox(checkRamUp, checkSameUserIter, infiniteIter)

	// Pole wprowadzania dla nazwy
	nameGroupEntry := widget.NewEntry()
	nameContainer := container.NewVBox(widget.NewLabel("Nazwa"), nameGroupEntry)
	commentsGroupEntry := widget.NewEntry()
	commentsContainer := container.NewVBox(widget.NewLabel("Komentarz"), commentsGroupEntry)
	actionGroupEntry := widget.NewEntry()
	actionContainer := container.NewVBox(widget.NewLabel("Akcja"), actionGroupEntry)
	entryContainer := container.NewVBox(nameContainer, commentsContainer, actionContainer)

	// Duration Container Lopps etc.
	itersGroupEntry := widget.NewEntry()
	iterContainer := container.NewVBox(widget.NewLabel("Liczba iteracji"), itersGroupEntry)
	lifeTimeEntry := widget.NewEntry()
	lifeTimeContainer := container.NewVBox(widget.NewLabel("Czas działania"), lifeTimeEntry)
	durationEntry := widget.NewEntry()
	durationContainer := container.NewVBox(widget.NewLabel("Duration"), durationEntry)

	iterMainContainer := container.NewVBox(iterContainer, lifeTimeContainer, durationContainer)

	saveBtn := widget.NewButton("Dodaj Grupę", func() {

	})

	GroupCreateContainer := container.NewVBox(checksContainer, entryContainer, iterMainContainer, saveBtn)

	return GroupCreateContainer
}
