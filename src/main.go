package main

import (
	_ "embed"
	guiController "gui-mini-ttracker/gui-controller"
	"gui-mini-ttracker/helpers"
	timer "gui-mini-ttracker/library/timer"
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

	t := timer.New()

	isStart := true

	controllerGUI.Common.StartButton.Connect("clicked", func() {
		if isStart {
			t.Start()
			go handleTimer(t, *controllerGUI)
		} else {
			t.Stop()
		}
		isStart = !isStart
	})

	gtk.Main()
}

func handleTimer(t timer.Timer, gui guiController.GUIInterface) {
	for seconds := range t.Seconds {
		gui.Common.Timer.SetText(helpers.ConvertSecondsToHumanTime(seconds))
	}
}
