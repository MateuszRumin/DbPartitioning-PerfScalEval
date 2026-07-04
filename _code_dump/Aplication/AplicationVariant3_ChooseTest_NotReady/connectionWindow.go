package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var connectionWindow fyne.Window

func setConnectionValues(a fyne.App) {

	connectionWindow = a.NewWindow("Połączenie")

	userEntry := widget.NewEntry()
	userEntry.SetText(Connection.User)
	hostEntry := widget.NewEntry()
	hostEntry.SetText(Connection.Host)
	passEntry := widget.NewEntry()
	passEntry.SetText(Connection.Password)
	portEntry := widget.NewEntry()
	portEntry.SetText(Connection.Port)
	dbEntry := widget.NewEntry()
	dbEntry.SetText(Connection.DatabaseName)

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

	saveBTN := widget.NewButton("Zapisz i wyjdź", func() {
		Connection.ModelInfo = "Non-Partitioned"
		Connection.User = userEntry.Text
		Connection.Host = hostEntry.Text
		Connection.Port = portEntry.Text
		Connection.Password = passEntry.Text
		Connection.DatabaseName = dbEntry.Text

		connectionWindow.Close()

	})

	connection := container.NewBorder(
		nil,     // góra
		saveBTN, // dół
		nil,     // lewo
		nil,     // prawo
		formContainer,
	)

	connectionWindow.SetContent(connection)

	connectionWindow.Resize(fyne.NewSize(400, 300))
	connectionWindow.Show()

}
