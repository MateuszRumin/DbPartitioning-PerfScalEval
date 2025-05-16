package applayout

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"perfscaleval/connectionssettings"
)

// func CreateConnectionsContent() fyne.CanvasObject {
// 	// Checkboxy na górze
// 	toggleSwitch := widget.NewCheck("ON/OFF", nil)
// 	configCheck := widget.NewCheck("Configuration", nil)
// 	partitionCheck := widget.NewCheck("Non-Partitioned", nil)
// 	checkContainer := container.NewHBox(
// 		toggleSwitch,
// 		configCheck,
// 		partitionCheck,
// 	)

// 	// Pola formularza
// 	userEntry := widget.NewEntry()
// 	hostEntry := widget.NewEntry()
// 	passEntry := widget.NewEntry()
// 	portEntry := widget.NewEntry()
// 	threadsEntry := widget.NewEntry()
// 	dbEntry := widget.NewEntry()

// 	// Układ w 2 kolumny
// 	form := container.NewGridWithColumns(2,
// 		container.NewVBox(
// 			widget.NewLabel("User:"),
// 			userEntry,
// 			widget.NewLabel("Password:"),
// 			passEntry,
// 			widget.NewLabel("Threads:"),
// 			threadsEntry,
// 		),
// 		container.NewVBox(
// 			widget.NewLabel("Host:"),
// 			hostEntry,
// 			widget.NewLabel("Port:"),
// 			portEntry,
// 			widget.NewLabel("Database Name:"),
// 			dbEntry,
// 		),
// 	)

// 	// Przycisk Start
// 	startBtn := widget.NewButton("Start", nil)
// 	btnContainer := container.NewCenter(startBtn)

// 	// Główny kontener
// 	return container.NewBorder(
// 		checkContainer, // góra
// 		btnContainer,   // dół
// 		nil,            // lewo
// 		nil,            // prawo
// 		form,           // środek
// 	)
// }

func createConnection1Layour() fyne.CanvasObject {
	border1 := canvas.NewRectangle(color.Transparent)
	border1.StrokeColor = color.White
	border1.StrokeWidth = 1

	toggleSwitchConnection := widget.NewCheck("ON/OFF Połączenie 1", func(checked bool) {
		connectionssettings.EnableConnection1 = checked
	})

	checkContainer := container.NewCenter(toggleSwitchConnection)

	userEntry := widget.NewEntry()
	hostEntry := widget.NewEntry()
	passEntry := widget.NewEntry()
	portEntry := widget.NewEntry()
	threadsEntry := widget.NewEntry()
	dbEntry := widget.NewEntry()

	formContainer := container.NewGridWithColumns(2,
		container.NewVBox(
			widget.NewLabel("User:"),
			userEntry,
			widget.NewLabel("Password:"),
			passEntry,
			widget.NewLabel("Threads:"),
			threadsEntry,
		),
		container.NewVBox(
			widget.NewLabel("Host:"),
			hostEntry,
			widget.NewLabel("Port:"),
			portEntry,
			widget.NewLabel("Database Name:"),
			dbEntry,
		),
	)

	formAkceppt := container.NewCenter(widget.NewLabel("Połaczenie 1"))

	mainContainer := container.NewBorder(
		checkContainer, // góra
		formAkceppt,    // dół
		nil,            // lewo
		nil,            // prawo
		formContainer,
	)

	connection1 := container.NewStack(border1, mainContainer)

	return connection1
}

func createConnection2Layour() fyne.CanvasObject {
	border1 := canvas.NewRectangle(color.Transparent)
	border1.StrokeColor = color.White
	border1.StrokeWidth = 1

	toggleSwitchConnection := widget.NewCheck("ON/OFF Połączenie 2", func(checked bool) {
		connectionssettings.EnableConnection1 = checked
	})

	checkContainer := container.NewCenter(toggleSwitchConnection)

	userEntry := widget.NewEntry()
	hostEntry := widget.NewEntry()
	passEntry := widget.NewEntry()
	portEntry := widget.NewEntry()
	threadsEntry := widget.NewEntry()
	dbEntry := widget.NewEntry()

	formContainer := container.NewGridWithColumns(2,
		container.NewVBox(
			widget.NewLabel("User:"),
			userEntry,
			widget.NewLabel("Password:"),
			passEntry,
			widget.NewLabel("Threads:"),
			threadsEntry,
		),
		container.NewVBox(
			widget.NewLabel("Host:"),
			hostEntry,
			widget.NewLabel("Port:"),
			portEntry,
			widget.NewLabel("Database Name:"),
			dbEntry,
		),
	)

	formAkceppt := container.NewCenter(widget.NewLabel("Połaczenie 2"))

	mainContainer := container.NewBorder(
		checkContainer, // góra
		formAkceppt,    // dół
		nil,            // lewo
		nil,            // prawo
		formContainer,
	)

	connection1 := container.NewStack(border1, mainContainer)

	return connection1
}

func CreateConnectionsContent() fyne.CanvasObject {
	// Tworzymy osobne obramowania dla każdego kontenera

	// Każdy kontener z własnym obramowaniem

	// Poprawne użycie GridWithRows - najpierw podajemy liczbę wierszy
	connections := container.NewGridWithRows(2, // dwa wiersze
		createConnection1Layour(),
		createConnection2Layour(),
	)

	return container.NewStack(connections) // dodajemy kontener maksymalizujący
}
