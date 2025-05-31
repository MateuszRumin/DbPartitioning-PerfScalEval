package applayout

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"perfscaleval/confmodel"
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

func CreateConnectionsContent() fyne.CanvasObject {
	////////////////
	border1 := canvas.NewRectangle(color.Transparent)
	border1.StrokeColor = color.White
	border1.StrokeWidth = 1

	toggleSwitchConnection := widget.NewCheck("ON/OFF Połączenie 1", func(checked bool) {

	})
	toggleSwitchConnection.SetChecked(confmodel.Connection1.Enable)
	checkContainer := container.NewCenter(toggleSwitchConnection)

	userEntry := widget.NewEntry()
	userEntry.SetText(confmodel.Connection1.User)
	hostEntry := widget.NewEntry()
	hostEntry.SetText(confmodel.Connection1.Host)
	passEntry := widget.NewEntry()
	passEntry.SetText(confmodel.Connection1.Password)
	portEntry := widget.NewEntry()
	portEntry.SetText(confmodel.Connection1.Port)
	dbEntry := widget.NewEntry()
	dbEntry.SetText(confmodel.Connection1.DatabaseName)

	Container := container.NewGridWithColumns(2,
		container.NewVBox(
			widget.NewLabel("User:"),
			userEntry,
			widget.NewLabel("Password:"),
			passEntry,
		),
		container.NewVBox(
			widget.NewLabel("Host:"),
			hostEntry,
			widget.NewLabel("Port:"),
			portEntry,
		),
	)
	dbnamecontainer := container.NewVBox(widget.NewLabel("Database Name:"), dbEntry)
	formContainer := container.NewVBox(Container, dbnamecontainer)

	mainContainer := container.NewBorder(
		checkContainer, // góra
		nil,            // dół
		nil,            // lewo
		nil,            // prawo
		formContainer,
	)

	connection1 := container.NewStack(border1, mainContainer)
	//////////

	border2 := canvas.NewRectangle(color.Transparent)
	border2.StrokeColor = color.White
	border2.StrokeWidth = 1

	toggleSwitchConnection2 := widget.NewCheck("ON/OFF Połączenie 2", func(checked bool) {

	})
	toggleSwitchConnection2.SetChecked(confmodel.Connection2.Enable)
	checkContainer2 := container.NewCenter(toggleSwitchConnection2)

	userEntry2 := widget.NewEntry()
	userEntry2.SetText(confmodel.Connection2.User)
	hostEntry2 := widget.NewEntry()
	hostEntry2.SetText(confmodel.Connection2.Host)
	passEntry2 := widget.NewEntry()
	passEntry2.SetText(confmodel.Connection2.Password)
	portEntry2 := widget.NewEntry()
	portEntry2.SetText(confmodel.Connection2.Port)
	dbEntry2 := widget.NewEntry()
	dbEntry2.SetText(confmodel.Connection2.DatabaseName)

	Container2 := container.NewGridWithColumns(2,
		container.NewVBox(
			widget.NewLabel("User:"),
			userEntry2,
			widget.NewLabel("Password:"),
			passEntry2,
		),
		container.NewVBox(
			widget.NewLabel("Host:"),
			hostEntry2,
			widget.NewLabel("Port:"),
			portEntry2,
		),
	)

	dbnamecontainer2 := container.NewVBox(widget.NewLabel("Database Name:"), dbEntry2)
	formContainer2 := container.NewVBox(Container2, dbnamecontainer2)

	mainContainer2 := container.NewBorder(
		checkContainer2, // góra
		nil,             // dół
		nil,             // lewo
		nil,             // prawo
		formContainer2,
	)

	connection2 := container.NewStack(border2, mainContainer2)
	///////
	btn_save := widget.NewButton("Zapisz", func() {
		confmodel.Connection1.Enable = toggleSwitchConnection.Checked
		confmodel.Connection1.ModelInfo = "Non-Partitioned"
		confmodel.Connection1.User = userEntry.Text
		confmodel.Connection1.Host = hostEntry.Text
		confmodel.Connection1.Port = portEntry.Text
		confmodel.Connection1.Password = passEntry.Text
		confmodel.Connection1.DatabaseName = dbEntry.Text

		confmodel.Connection2.Enable = toggleSwitchConnection2.Checked
		confmodel.Connection2.ModelInfo = "Partitioned"
		confmodel.Connection2.User = userEntry2.Text
		confmodel.Connection2.Host = hostEntry2.Text
		confmodel.Connection2.Port = portEntry2.Text
		confmodel.Connection2.Password = passEntry2.Text
		confmodel.Connection2.DatabaseName = dbEntry2.Text
	})

	connections := container.NewGridWithRows(2, // dwa wiersze
		connection1,
		connection2,
	)

	return container.NewBorder(nil, btn_save, nil, nil, connections) // dodajemy kontener maksymalizujący
}
