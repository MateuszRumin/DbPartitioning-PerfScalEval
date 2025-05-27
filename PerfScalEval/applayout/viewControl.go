package applayout

import "fyne.io/fyne/v2"

var (
	currentMainContent *fyne.Container
)

func InitViewManager(content *fyne.Container) {
	currentMainContent = content
}

func SwitchView(content fyne.CanvasObject) {
	if currentMainContent != nil {
		currentMainContent.Objects = []fyne.CanvasObject{content}
		currentMainContent.Refresh()
	}
}
