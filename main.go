package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	mainWindow  fyne.Window
	mainContent *fyne.Container
)

func main() {
	a := app.New()
	mainWindow = a.NewWindow("Perf-Scal-eval")
	mainWindow.Resize(fyne.NewSize(300, 300))

	initMainContent(a)

	mainWindow.Show()
	a.Run()
}

func initMainContent(a fyne.App) {
	mainContent = container.NewStack()

	mainContent.Add(mainLayout(a))

	topMenu := topMenu(a) // Teraz bez callbacka
	content := container.NewBorder(topMenu, nil, nil, nil, mainContent)

	// Ładowanie domyślnego widoku
	// applayout.SwitchView(applayout.CreateConnectionsContent())

	mainWindow.SetContent(content)
}

func topMenu(a fyne.App) fyne.CanvasObject {

	connectionsBtn := widget.NewButton("Ustawienia danych połączenia z bazą", func() {
		setConnectionValues(a)

	})
	connectionsBtn.Resize(fyne.NewSize(220, 50))

	topMenu := container.NewHBox(
		connectionsBtn,
	)

	return container.NewCenter(topMenu)
}

func mainLayout(a fyne.App) fyne.CanvasObject {

	connectionsBtn := widget.NewButton("Test1", func() {
		test1PopUP(a)

	})

	container1 := container.NewVBox(
		connectionsBtn,
	)
	container2 := container.NewVBox(
		widget.NewLabel("Test2"),
	)

	mainContent := container.NewHBox(
		container1,
		container2,
	)

	return container.NewCenter(mainContent)
}
