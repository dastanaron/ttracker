package helpers

import (
	guiController "gui-mini-ttracker/gui-controller"
	"log"
)

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func CheckGUIError(gui *guiController.GUIInterface, title string, err error) {
	if err != nil {
		gui.ErrorDialog.ErrorDialog.ShowAll()
		gui.ErrorDialog.ErrorDialog.SetMarkup(title)
		gui.ErrorDialog.ErrorDialog.FormatSecondaryMarkup("%s", err.Error())
	}
}
