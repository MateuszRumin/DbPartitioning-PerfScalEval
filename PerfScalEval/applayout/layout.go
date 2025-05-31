package applayout

import (
	"encoding/json"
	"fmt"
	"io"
	"perfscaleval/confmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Globalna zmienna do przechowywania referencji

func load() {
	// Open file dialog to choose file to load
	openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
		if err != nil {
			fmt.Println("Error selecting file:", err)
			return
		}
		if read == nil {
			// User cancelled
			return
		}
		defer read.Close()

		// Read the file content
		jsonData, err := io.ReadAll(read)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Define the config structure to unmarshal into
		type Config struct {
			Connections struct {
				Connection1 confmodel.ConnectionSetting `json:"connection1"`
				Connection2 confmodel.ConnectionSetting `json:"connection2"`
			} `json:"connections"`
			TestPlans    []confmodel.TestPlan `json:"test_plans"`
			CurrentState struct {
				CreateMode   bool `json:"create_mode"`
				GroupExist   bool `json:"group_exist"`
				ChooseLayOut int  `json:"choose_layout"`
			} `json:"current_state"`
		}

		var config Config

		// Unmarshal the JSON
		if err := json.Unmarshal(jsonData, &config); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		// Update the global variables
		confmodel.Connection1 = config.Connections.Connection1
		confmodel.Connection2 = config.Connections.Connection2
		confmodel.Plan = config.TestPlans
		confmodel.CreateMode = config.CurrentState.CreateMode
		confmodel.GroupExist = config.CurrentState.GroupExist
		confmodel.ChooseLayOut = config.CurrentState.ChooseLayOut

		// If there are test plans, set the first one as current
		if len(confmodel.Plan) > 0 {
			confmodel.CurrentPlan = &confmodel.Plan[0]
		}

		// Show success message
		dialog.ShowInformation("Sukces", "Plan testowy został wczytany", fyne.CurrentApp().Driver().AllWindows()[0])

		// Here you might want to add code to refresh your UI to reflect the loaded data
		// refreshUI() or similar function if you have one

	}, fyne.CurrentApp().Driver().AllWindows()[0])

	// Set file filter and show dialog
	openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	openDialog.Show()

}

func save() {
	// Create a struct to hold all the configuration
	type Config struct {
		Connections struct {
			Connection1 confmodel.ConnectionSetting `json:"connection1"`
			Connection2 confmodel.ConnectionSetting `json:"connection2"`
		} `json:"connections"`
		TestPlans    []confmodel.TestPlan `json:"test_plans"`
		CurrentState struct {
			CreateMode   bool `json:"create_mode"`
			GroupExist   bool `json:"group_exist"`
			ChooseLayOut int  `json:"choose_layout"`
		} `json:"current_state"`
	}

	// Populate the config struct
	var config Config
	config.Connections.Connection1 = confmodel.Connection1
	config.Connections.Connection2 = confmodel.Connection2
	config.TestPlans = confmodel.Plan
	config.CurrentState.CreateMode = confmodel.CreateMode
	config.CurrentState.GroupExist = confmodel.GroupExist
	config.CurrentState.ChooseLayOut = confmodel.ChooseLayOut

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Open file dialog to choose save location
	saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
		if err != nil {
			fmt.Println("Error selecting file:", err)
			return
		}
		if write == nil {
			// User cancelled
			return
		}
		defer write.Close()

		// Write the JSON to file
		_, err = write.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		// Show success message
		dialog.ShowInformation("Sukces", "Plan testowy został zapisany", fyne.CurrentApp().Driver().AllWindows()[0])
	}, fyne.CurrentApp().Driver().AllWindows()[0])

	// Set file filter and show dialog
	saveDialog.SetFileName("test_plan.json")
	saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	saveDialog.Show()

}

func LayOut() fyne.CanvasObject {
	savePlan := widget.NewButton("Zapisz", func() { save() })
	loadPlan := widget.NewButton("Wczytaj", func() { load() })

	saveings := container.NewVBox(savePlan, loadPlan)

	connectionsBtn := widget.NewButton("Połączenia", func() {
		SwitchView(CreateConnectionsContent())
	})

	testPlanBtn := widget.NewButton("Test Plan", func() {
		SwitchView(CreateTestPlanContent())
	})
	resultsBtn := widget.NewButton("Wyniki", func() {
		SwitchView(CreateResultsContent())
	})

	return container.NewGridWithColumns(4, connectionsBtn, testPlanBtn, resultsBtn, saveings)
}
