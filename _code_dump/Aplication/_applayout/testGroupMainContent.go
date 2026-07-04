package applayout

import (
	"fmt"
	"perfscaleval/confmodel"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GroupCreate() fyne.CanvasObject {

	// Pole wprowadzania dla nazwy
	nameGroupEntry := widget.NewEntry()
	nameContainer := container.NewVBox(widget.NewLabel("Nazwa"), nameGroupEntry)
	commentsGroupEntry := widget.NewEntry()
	commentsContainer := container.NewVBox(widget.NewLabel("Komentarz"), commentsGroupEntry)
	actionGroupEntry := widget.NewEntry()
	actionContainer := container.NewVBox(widget.NewLabel("Akcja"), actionGroupEntry)
	threadsEntry := widget.NewEntry()
	threadsContainer := container.NewVBox(widget.NewLabel("Liczba połączeń"), threadsEntry)
	entryContainer := container.NewVBox(nameContainer, commentsContainer, actionContainer, threadsContainer)

	// Duration Container Lopps etc.
	itersGroupEntry := widget.NewEntry()
	iterContainer := container.NewVBox(widget.NewLabel("Liczba iteracji"), itersGroupEntry)
	lifeTimeEntry := widget.NewEntry()
	lifeTimeContainer := container.NewVBox(widget.NewLabel("Czas działania"), lifeTimeEntry)
	durationEntry := widget.NewEntry()
	durationContainer := container.NewVBox(widget.NewLabel("Duration"), durationEntry)
	rumpUpEntry := widget.NewEntry()
	rampupContainer := container.NewVBox(widget.NewLabel("rampUp"), rumpUpEntry)

	iterMainContainer := container.NewVBox(iterContainer, lifeTimeContainer, durationContainer, rampupContainer)

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
	checksContainer := container.NewHBox(checkSameUserIter, infiniteIter)

	saveBtn := widget.NewButton("Dodaj Grupę", func() {
		if nameGroupEntry.Text == "" || commentsGroupEntry.Text == "" || actionGroupEntry.Text == "" {
			dialog.ShowInformation("Błąd", "Uzupełnij pola", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		threadNum, err := strconv.Atoi(threadsEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		loopNum, err := strconv.Atoi(itersGroupEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		lifeTimeNum, err := strconv.Atoi(lifeTimeEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		durationNum, err := strconv.Atoi(durationEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		rumpUpNum, err := strconv.Atoi(rumpUpEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		newGroup := confmodel.TestGroup{
			Name:         nameGroupEntry.Text,
			Comments:     commentsGroupEntry.Text,
			Action:       actionGroupEntry.Text,
			ThreadNumber: threadNum,
			RampUp:       rumpUpNum,
			SameUserIter: checkSameUserIter.Checked,
			Loops:        loopNum,
			Infinite:     infiniteIter.Checked,
			LifeTime:     lifeTimeNum,
			Duration:     durationNum,
			Steps:        []confmodel.TestStep{},
		}
		confmodel.CurrentPlan.Group = append(confmodel.CurrentPlan.Group, newGroup)
		confmodel.CurrentGroup = &confmodel.CurrentPlan.Group[len(confmodel.Plan)-1]
		confmodel.ChooseLayOut = 1
		confmodel.CreateMode = false
		SwitchView(CreateTestPlanContent())
	})

	GroupCreateContainer := container.NewVBox(checksContainer, entryContainer, iterMainContainer, saveBtn)

	return GroupCreateContainer
}

func GroupEdit() fyne.CanvasObject {

	// Pole wprowadzania dla nazwy
	nameGroupEntry := widget.NewEntry()
	nameGroupEntry.SetText(confmodel.CurrentGroup.Name)
	nameContainer := container.NewVBox(widget.NewLabel("Nazwa"), nameGroupEntry)

	commentsGroupEntry := widget.NewEntry()
	commentsGroupEntry.SetText(confmodel.CurrentGroup.Comments)
	commentsContainer := container.NewVBox(widget.NewLabel("Komentarz"), commentsGroupEntry)

	actionGroupEntry := widget.NewEntry()
	actionGroupEntry.SetText(confmodel.CurrentGroup.Action)
	actionContainer := container.NewVBox(widget.NewLabel("Akcja"), actionGroupEntry)

	threadsEntry := widget.NewEntry()
	threadsEntry.SetText(fmt.Sprintf("%d", confmodel.CurrentGroup.ThreadNumber))
	threadsContainer := container.NewVBox(widget.NewLabel("Liczba połączeń"), threadsEntry)

	entryContainer := container.NewVBox(nameContainer, commentsContainer, actionContainer, threadsContainer)

	// Duration Container Lopps etc.
	itersGroupEntry := widget.NewEntry()
	itersGroupEntry.SetText(fmt.Sprintf("%d", confmodel.CurrentGroup.Loops))
	iterContainer := container.NewVBox(widget.NewLabel("Liczba iteracji"), itersGroupEntry)

	lifeTimeEntry := widget.NewEntry()
	lifeTimeEntry.SetText(fmt.Sprintf("%d", confmodel.CurrentGroup.LifeTime))
	lifeTimeContainer := container.NewVBox(widget.NewLabel("Czas działania"), lifeTimeEntry)

	durationEntry := widget.NewEntry()
	durationEntry.SetText(fmt.Sprintf("%d", confmodel.CurrentGroup.Duration))
	durationContainer := container.NewVBox(widget.NewLabel("Duration"), durationEntry)

	rumpUpEntry := widget.NewEntry()
	rumpUpEntry.SetText(fmt.Sprintf("%d", confmodel.CurrentGroup.RampUp))
	rampupContainer := container.NewVBox(widget.NewLabel("rampUp"), rumpUpEntry)

	iterMainContainer := container.NewVBox(iterContainer, lifeTimeContainer, durationContainer, rampupContainer)

	checkSameUserIter := widget.NewCheck("Same User Iter", func(checked bool) {
		if checked {
			println("Checkbox is checked")
		} else {
			println("Checkbox is unchecked")
		}
	})
	checkSameUserIter.SetChecked(confmodel.CurrentGroup.SameUserIter)
	infiniteIter := widget.NewCheck("Infinite Loop", func(checked bool) {
		if checked {
			println("Checkbox is checked")
		} else {
			println("Checkbox is unchecked")
		}
	})
	infiniteIter.SetChecked(confmodel.CurrentGroup.Infinite)
	checksContainer := container.NewHBox(checkSameUserIter, infiniteIter)

	modBtn := widget.NewButton("Modyfikuj", func() {
		if nameGroupEntry.Text == "" || commentsGroupEntry.Text == "" || actionGroupEntry.Text == "" {
			dialog.ShowInformation("Błąd", "Uzupełnij pola", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		threadNum, err := strconv.Atoi(threadsEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		loopNum, err := strconv.Atoi(itersGroupEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		lifeTimeNum, err := strconv.Atoi(lifeTimeEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		durationNum, err := strconv.Atoi(durationEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		rumpUpNum, err := strconv.Atoi(rumpUpEntry.Text)
		if err != nil {
			dialog.ShowInformation("Błąd", "Nieprawidłowa liczba wątków", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		confmodel.CurrentGroup.Name = nameGroupEntry.Text
		confmodel.CurrentGroup.Comments = commentsGroupEntry.Text
		confmodel.CurrentGroup.Action = actionGroupEntry.Text
		confmodel.CurrentGroup.ThreadNumber = threadNum
		confmodel.CurrentGroup.RampUp = rumpUpNum
		confmodel.CurrentGroup.SameUserIter = checkSameUserIter.Checked
		confmodel.CurrentGroup.Loops = loopNum
		confmodel.CurrentGroup.Infinite = infiniteIter.Checked
		confmodel.CurrentGroup.LifeTime = lifeTimeNum
		confmodel.CurrentGroup.Duration = durationNum
		confmodel.ChooseLayOut = 1
		confmodel.CreateMode = false
		SwitchView(CreateTestPlanContent())
	})

	GroupCreateContainer := container.NewVBox(checksContainer, entryContainer, iterMainContainer, modBtn)

	return GroupCreateContainer
}
