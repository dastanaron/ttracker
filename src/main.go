package main

import (
	_ "embed"
	guiController "gui-mini-ttracker/gui-controller"
	"gui-mini-ttracker/helpers"
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

//go:embed main.glade
var gladeInterface string

//go:embed icon.png
var icon []byte

func main() {
	gtk.Init(nil)

	pixBuf, err := gdk.PixbufNewFromBytesOnly(icon)
	helpers.CheckError("Error", err)

	gui, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = gui.AddFromString(gladeInterface)
	helpers.CheckError("Error", err)

	controllerGUI, err := guiController.NewGUIController(gui)
	helpers.CheckError("Error", err)

	controllerGUI.Common.MainWindow.SetTitle("MINI tTracker")
	controllerGUI.Common.MainWindow.SetIcon(pixBuf)
	controllerGUI.Common.MainWindow.ToWindow().SetIcon(pixBuf)
	controllerGUI.Common.MainWindow.Connect("destroy", func() {
		gtk.MainQuit()
	})

	controllerGUI.ErrorDialog.CloseButton.Connect("clicked", func() {
		controllerGUI.ErrorDialog.ErrorDialog.Hide()
	})

	controllerGUI.Common.MainWindow.ShowAll()

	gtk.Main()
}
